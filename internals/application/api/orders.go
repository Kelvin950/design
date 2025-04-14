package api

import "github.com/kelvin950/desing/internals/application/domain"


func(a Api)CreateOrders(orderDets  *[]domain.OrderDetail, userId int)(domain.Orders ,error){
  


	TotalQuantity := 0 
	totalprice :=0.0

 
	for _ , order_detail :=range *orderDets{
        
		TotalQuantity += order_detail.Quantity
		totalprice+= order_detail.Price 
	}

	var id =  uint(userId)
	var order =  domain.Orders{
		UserrID:&id ,
		TotalQuantity: TotalQuantity,
		TotalPrice: totalprice, 
		OrderDetails: *orderDets,
	}

	err:= a.DB.CreateOrders(&order) 
  if err!=nil{
	return  domain.Orders{} ,err 
  }


	return order , nil
}

func (a Api)GetOrders(orders *[]domain.Orders)error{


  err:=	a.DB.GetOrders(orders)

  if err!=nil{
	return err
  }


  return nil
}



func(a Api)GetUserOrder(order *domain.Orders)error{
	
	err:= a.DB.GetUserOrders(order)

	if err!=nil{
		return err 
	}

	return nil
}