package db

import (
	"errors"
	"fmt"
	"log"


	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserrID *uint `gorm:"column:user_id"`
	TotalQuantity int `gorm:"column:total_quantity"` 
	TotalPrice float64 `gorm:"column:total_price"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}



func(d DB)CreateOrders(order *domain.Orders)error{


	orderDets := []OrderDetail{}

	log.Print(order.OrderDetails)
	for _ ,r:= range order.OrderDetails{
		orderDets = append(orderDets, OrderDetail{
			ProductID: r.ProductID,
			Price: r.Price,
			Quantity: r.Quantity,
		})
	}

	var newOrder = Order{
		UserrID: order.UserrID,
		TotalQuantity: order.TotalQuantity,
		TotalPrice: order.TotalPrice, 
		OrderDetails:orderDets,
		
	} 

	

	 result:= d.db.Create(&newOrder)
	   


 if result.Error!=nil{
		return result.Error
	 }
	 
	 fmt.Println(newOrder)
	 

	 order.Id = newOrder.ID
	 order.Createdat=  newOrder.CreatedAt
	 order.Updatedat= newOrder.UpdatedAt
	 
	 for _ ,r:= range newOrder.OrderDetails{
		 order.OrderDetails = append(order.OrderDetails, domain.OrderDetail{
			ProductID: r.ProductID,
			OrderID:  r.OrderID,
			Id: r.ID ,
			Price: r.Price,
			Quantity: r.Quantity,
			Createdat: r.CreatedAt,
			Updatedat: r.UpdatedAt,
		 })
	 }
	
	 return nil
}


func (d DB)GetOrders(orders  *[]domain.Orders)error{

	var  ods []Order 

	result := d.db.Preload("OrderDetails.Product").Find(&ods) 

	if result.Error!=nil{
		return result.Error
	}
 

	for _ , r:= range ods{
	
	order := domain.Orders{
			Id: r.ID,
			UserrID:r.UserrID,
			Createdat: r.CreatedAt,
			Updatedat: r.UpdatedAt,
			TotalQuantity: r.TotalQuantity,
			TotalPrice: r.TotalPrice, 
			

		}
	for _ , v:= range r.OrderDetails{
		order.OrderDetails =  append(order.OrderDetails, domain.OrderDetail{
			ProductID: v.ProductID,
			Id: v.ID,
			Price: v.Price,
		 Product: domain.Product{
			Id: v.Product.ID, 
			Name: v.Product.Name,
			Color: v.Product.Color,
			Sizes: v.Product.Sizes, 
			Quantity: v.Product.Quantity,
			MinPrice: v.Product.MinPrice,
			MaxPrice: v.Product.MaxPrice,
			BrandID: v.Product.BrandID, 
			Createdat: v.Product.CreatedAt,
			Updatedat: v.Product.UpdatedAt,
		 },
			Quantity: v.Quantity,
			Createdat: v.CreatedAt,
			Updatedat: v.UpdatedAt,
			OrderID: v.OrderID,
		})
	}
		*orders =  append(*orders,order)


	}
	return nil
}

func(d DB)GetUserOrders(order  *domain.Orders)error{
	 
   var ord Order 

  result :=  d.db.Preload("OrderDetails.Product").First(&ord , &Order{UserrID: order.UserrID})  

  if result.Error !=nil{
	if result.Error !=nil{
		 if errors.Is(result.Error , gorm.ErrRecordNotFound){
			return nil
		 }

		 return result.Error
	}
	return result.Error 
  }

var order_details []domain.OrderDetail
  for _ , orderDet := range ord.OrderDetails{

	order_details =  append(order_details, domain.OrderDetail{
		ProductID: orderDet.ProductID,
		OrderID: orderDet.OrderID,
		 Product: domain.Product{
			Id: orderDet.Product.ID, 
			Name: orderDet.Product.Name,
			Color: orderDet.Product.Color,
			Sizes: orderDet.Product.Sizes, 
			Quantity: orderDet.Product.Quantity,
			MinPrice: orderDet.Product.MinPrice,
			MaxPrice: orderDet.Product.MaxPrice,
			BrandID: orderDet.Product.BrandID, 
			Createdat: orderDet.Product.CreatedAt,
			Updatedat: orderDet.Product.UpdatedAt,
		 },
		Id: orderDet.ID,
		Price: orderDet.Price,
		Quantity: orderDet.Quantity,
		Createdat: order.Createdat,
		Updatedat: orderDet.UpdatedAt, 
	})
  }
 order.TotalQuantity =  ord.TotalQuantity 
 order.TotalPrice =  ord.TotalPrice
   order.OrderDetails = order_details
   order.Createdat =  ord.CreatedAt 
   order.Id =  ord.ID 
   order.Updatedat=  ord.UpdatedAt

	return nil
}