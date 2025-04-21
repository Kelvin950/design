package db

import (
	

	"github.com/kelvin950/desing/internals/application/domain"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name string 
	Color  pq.Int64Array `gorm:"type:integer[]"`
	Sizes  pq.Int64Array `gorm:"type:integer[]"` 
	Quantity int  
	MinPrice float64  `gorm:"column:min_price"`
	MaxPrice float64  `gorm:"column:max_price"`
	Typpe Typpe  `gorm:"foreignKey:TyppeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Brand Brand `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TyppeID  *uint  `gorm:"column:type_id"`
	BrandID *uint  `gorm:"column:brand_id"`
	ProductDetails []Product_Detail `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderDetails []OrderDetail
	Wishlist []Wishlist `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (d DB) CreateProduct(product *domain.Product)error{ 

	 var newProduct = Product{
		Name: product.Name, 
		Color: product.Color,
		Sizes: product.Sizes,
		Quantity: product.Quantity, 
		MinPrice: product.MinPrice,
		MaxPrice: product.MaxPrice,
		BrandID: product.BrandID,
		TyppeID: product.TyppeID,
	 }

   result := d.db.Save(&newProduct)
   if result.Error !=nil{
	return result.Error 
   }
 

   product.Id =  newProduct.ID 
   product.Createdat =  newProduct.CreatedAt
    product.Updatedat =  newProduct.UpdatedAt
   return nil 
   
}


func (d DB)GetProducts(products *[]domain.Product)error{
  
	var prod  []Product 

	 result :=d.db.Preload("Typpe").Preload("Brand").Find(&prod) 
	 if result.Error !=nil{
		return result.Error
	 }

	 for _ , p := range prod{
     *products =  append(*products, domain.Product{
		Name: p.Name,
		Id: p.ID,
		BrandID: p.BrandID,
		TyppeID: p.TyppeID,
		Brand: domain.Brand{
		 Name: p.Brand.Name,
		 Id: p.Brand.ID, 
		 Createdat: p.Brand.CreatedAt,
		 Updatedat: p.Brand.UpdatedAt,
		},
		Typpe: domain.Typpe{
			Id: p.Typpe.ID,
			Name: p.Typpe.Name,
			Createdat: p.Typpe.CreatedAt,
			Updatedat: p.Typpe.UpdatedAt ,
			AudienceID: p.Typpe.AudienceID,
			CategoryID: p.Typpe.CategoryID,
			UnitOfMeasurement: p.Typpe.UnitOfMeasurement,

		},
		Createdat: p.CreatedAt,
		Updatedat: p.UpdatedAt,
		MinPrice: p.MinPrice,
		MaxPrice: p.MaxPrice,
		Quantity: p.Quantity, 
		Color: p.Color, 
		Sizes: p.Sizes,
	 })

	 } 


 
	 return nil 

}

func(d DB)GetProduct(product *domain.Product)error{
	var prod  Product 

	result := d.db.Find(&prod , product.Id) 
	if result.Error!=nil {
		return result.Error
	}

	product.Name =  prod.Name
	product.Color =  prod.Color 
	product.BrandID =  prod.BrandID 
	product.TyppeID =  prod.TyppeID 
	product.MaxPrice =prod.MaxPrice 
	product.MinPrice = prod.MinPrice
	product.Sizes =  prod.Sizes 
	product.Createdat =  prod.CreatedAt 
	product.Updatedat =  prod.UpdatedAt 
	return  nil 	
}