package ports

import "github.com/kelvin950/desing/internals/application/domain"

type DB interface{
	CreateUser(user *domain.User)error 
	GetUsers(*[]domain.User)(error) 
	GetUser(user *domain.User)error
	GetUserByUserName(user *domain.User)error 
	GetUserBYEmail(user *domain.User)error
	CreateCategory(category *domain.Category)error
	CreateAudience(audience  *domain.Audience)error
	CreateTypes(t *domain.Typpe)error
	CreateColor(color *domain.Color)error
	CreateProduct(product *domain.Product)error
	GetProducts(products *[]domain.Product)error
	GetProduct(product *domain.Product)error
	CreateProduct_Detail(product_detail *domain.Product_Detail)error
	GetProductDetails(product_details  *[]domain.Product_Detail)error
	CreateOrders(order *domain.Orders)error
	GetOrders(orders  *[]domain.Orders)error
	GetUserOrders(order  *domain.Orders)error
	GetUserByFirebaseUID(user *domain.User)error
	 CreateWishlist(wishlist []domain.Wishlist)error 
	GetWishlistByUserID(wishlist []domain.Wishlist)([]domain.Wishlist , error)
}

