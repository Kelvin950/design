package ports

import "github.com/kelvin950/desing/internals/application/domain"

type Api interface{
	LoginUser(cred string , aud string)( domain.User,string , error)
	 CreateProduct(product *domain.Product)error
	 CreateUser(user *domain.User)error
	  GetProducts(products *[]domain.Product)error
	  GetProduct(product *domain.Product)error
	  CreateProduct_Detail(product_detail *domain.Product_Detail)error
	  GetProductDetails(product_details *[]domain.Product_Detail)([]domain.Product_Detail_Product ,error)
	  CreateOrders(orderDets  *[]domain.OrderDetail, userId int)(domain.Orders ,error)
	  GetOrders(orders *[]domain.Orders)error
	  GetUserOrder(order *domain.Orders)error
	  SigInUser(user *domain.User)(error)
}