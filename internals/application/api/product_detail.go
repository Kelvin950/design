package api

import (
	"fmt"

	"github.com/kelvin950/desing/internals/application/domain"
)

func (a Api) CreateProduct_Detail(product_detail *domain.Product_Detail)error{

	 err:=a.DB.CreateProduct_Detail(product_detail)
	 if err!=nil{
		return err
	 }

	 return nil 

}


func(a Api)GetProductDetails(product_details *[]domain.Product_Detail)([]domain.Product_Detail_Product , error){

	err := a.DB.GetProductDetails(product_details) 
	if err!=nil{
		return []domain.Product_Detail_Product{} ,  err 
	}
 

	var products =   make(map[uint]domain.Product_Detail_Product)

	for _ , product_detail := range *product_details {

		 
	   det,ok :=products[product_detail.Product.Id] 

	   if !ok {
			det =domain.Product_Detail_Product{ 
			Name:  product_detail.Product.Name ,
			CreatedAt:product_detail.Product.Createdat ,
			ID: product_detail.Product.Id,
			UpdatedAt:product_detail.Product.Updatedat ,
			Quantity:product_detail.Product.Quantity  ,
			Price: product_detail.Product.MinPrice ,
			ImageUrl: []string{product_detail.Product_Images.Image1 , 
			product_detail.Product_Images.Image2 , 
			product_detail.Product_Images.Image3 , product_detail.Product_Images.Image4} ,
			Colors: []domain.Product_Detail_Colors{
				{
					Name:product_detail.Color.Name ,
					ID:product_detail.Color.Id  ,
					Code: product_detail.Color.Code,
					Sizes: []domain.Product_Detail_Size{
						 {
				Name: product_detail.Sizee.Name, 
				Price: product_detail.Price,
				Quantity: product_detail.Quantity,
			} ,
					} ,
				} ,
			} ,
		}
	
	   }else{ 
		
		 var index = -1
		for  i , r:= range  det.Colors{
			if r.ID ==  product_detail.Color.Id {
			
				index = i
				break
			}
		}

		if index > -1 {
		
			det.Colors[index].Sizes =  append(det.Colors[index].Sizes,  domain.Product_Detail_Size{
				Name: product_detail.Sizee.Name, 
				Price: product_detail.Price,
				Quantity: product_detail.Quantity,
			})
		}else {
   fmt.Println(product_detail.Color.Name )
			det.Colors=append(det.Colors, domain.Product_Detail_Colors{
				Name:product_detail.Color.Name ,
					ID:product_detail.Color.Id  ,
					Code: product_detail.Color.Code,
					Sizes:[]domain.Product_Detail_Size{
						 {
				Name: product_detail.Sizee.Name, 
				Price: product_detail.Price,
				Quantity: product_detail.Quantity,
			} ,
					} ,
			})

			fmt.Println(det.Colors)
		}

			
	   }
	   	products[product_detail.Product.Id] =det
	}

	var returnProd []domain.Product_Detail_Product

	for _ ,r:= range products {
		returnProd = append(returnProd, r) 
	}
	return returnProd , nil
}