package api

import "github.com/kelvin950/desing/internals/application/domain"

func (a Api) CreateProduct(product *domain.Product)error{
 
	 
	err:= a.DB.CreateProduct(product) 

	 if err!=nil{
		return err
	 }


	 return nil 
}

func (a Api) GetProducts(products *[]domain.Product)error{
 

	err:= a.DB.GetProducts(products) 
	if err!=nil{
		return err 
	}

	return   nil
	
}



func(a Api)GetProduct(product *domain.Product)error{

	err:=a.DB.GetProduct(product)
	if err!=nil{
		return err 
	}

	return nil
}