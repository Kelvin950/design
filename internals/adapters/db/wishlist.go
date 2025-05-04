package db

import (
	"errors"

	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Wishlist struct {
	gorm.Model
   UserrID *uint `gorm:"column:user_id;uniqueIndex:idx_user_product"`
   User Userr  `gorm:"foreignKey:UserrID;constraint:OnUpdate:CasCADE,OnDelete:SET NULL;"`
   ProductID *uint `gorm:"column:product_id;uniqueIndex:idx_user_product"`
   Product  Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (d DB) CreateWishlist(wishlist *[]domain.Wishlist)error { 
var newWishlists  []Wishlist 


  for _ ,r:= range *wishlist{

	newWishlists = append(newWishlists, Wishlist{
		UserrID: r.UserID , 
		ProductID:  &r.ProductId,
	})

  }

 result := d.db.Save(&newWishlists) 

 if result.Error !=nil{
	return result.Error 
 }
 
 wishlist = &[]domain.Wishlist{}
for  _ , r:= range newWishlists{
*wishlist = append(*wishlist, domain.Wishlist{
	UserID: r.UserrID, 
	ID: r.ID,
	Createdat: r.CreatedAt,
	Updatedat: r.UpdatedAt, 
	Products: domain.Product{
		Id: *r.ProductID,
	},
})
}

  return nil
}

func (d DB) GetWishlistByUserID(wishlist []domain.Wishlist)([]domain.Wishlist , error)  {

	var userWishlist []Wishlist
     
	result := d.db.Preload("Product").Find(&userWishlist ,Wishlist{
		UserrID:(wishlist)[0].UserID	})

    if result.Error!=nil{

		  if errors.Is(gorm.ErrRecordNotFound ,result.Error){
				return []domain.Wishlist{} , nil
		  }
		return nil ,  result.Error
	}

	wishlist = []domain.Wishlist{}

	for _,r:=range userWishlist{

		wishlist =  append(wishlist, domain.Wishlist{
			UserID: r.UserrID,
			ID:  r.ID, 
			Createdat: r.CreatedAt , 
			Updatedat: r.UpdatedAt, 
			Products:domain.Product{
				Id: *r.ProductID, 
				Name: r.Product.Name, 
				BrandID: r.Product.BrandID,
				TyppeID: r.Product.TyppeID,
				Color: r.Product.Color,
				Sizes: r.Product.Sizes,
				Quantity: r.Product.Quantity,
				MinPrice: r.Product.MinPrice,
				MaxPrice: r.Product.MaxPrice,
				Createdat: r.Product.CreatedAt,
				Updatedat: r.Product.UpdatedAt,

			},
		})
	}

	

	return wishlist ,  nil
}