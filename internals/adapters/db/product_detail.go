package db

import (
	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Product_Detail struct {
	gorm.Model 
	ProductID *uint  `gorm:"column:product_id;uniqueIndex:product_detail_product_id_color_id_sizes_id_key"`
	Product Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Color Color `gorm:"foreignKey:ColorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sizee Sizee `gorm:"foreignKey:SizeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ColorID *uint  `gorm:"column:color_id;uniqueIndex:product_detail_product_id_color_id_sizes_id_key"`
	SizeeID *uint  `gorm:"column:sizes_id;uniqueIndex:product_detail_product_id_color_id_sizes_id_key"`
	Quantity int 
	Price float64 
	ProductImageID *uint  `gorm:"column:product_image_id"` 
	Product_Images Product_Images `gorm:"foreignKey:ProductImageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


func(d DB) CreateProduct_Detail(product_detail *domain.Product_Detail)error{

   var newProductDetail =  Product_Detail{
	 ProductID: product_detail.ProductID,
	 ColorID: product_detail.ColorID,
	 SizeeID: product_detail.SizeeID,
	 Quantity: product_detail.Quantity,
	 Price: product_detail.Price,
	 ProductImageID: product_detail.ProductImageID,
   }
   
   result :=d.db.Save(&newProductDetail) 

    if result.Error !=nil{
		return result.Error
	}
 

	product_detail.Id = newProductDetail.ID
	product_detail.Createdat = newProductDetail.CreatedAt
	product_detail.Updatedat =  newProductDetail.UpdatedAt
	return nil 

}


func(d DB)GetProductDetails(product_details  *[]domain.Product_Detail)error{
	 
	var product_dets []Product_Detail 
	result := d.db.Preload("Color").Preload("Product").Preload("Sizee").Preload("Product_Images").Find(&product_dets) 
	 
	if result.Error !=nil{
		return result.Error
	}

	for _ ,dd:= range product_dets {
		*product_details =  append(*product_details, domain.Product_Detail{
			ProductID: dd.ProductID,
			Id: dd.ID,
			Product: domain.Product{
				Name: dd.Product.Name,
				Id: dd.Product.ID,
				Quantity: dd.Product.Quantity,
				BrandID: dd.Product.BrandID,
				TyppeID: dd.Product.TyppeID, 
				Createdat: dd.Product.CreatedAt,
				Updatedat: dd.Product.UpdatedAt,
			},
			Color: domain.Color{
				Name: dd.Color.Name,
				Code: dd.Color.Code,
				Id:dd.Color.ID ,
				Createdat: dd.Color.CreatedAt,
				Updatedat: dd.Color.UpdatedAt,
			},
			Sizee: domain.Sizee{
				Name: dd.Sizee.Name,
				Id: dd.Sizee.ID,
				Createdat: dd.Sizee.CreatedAt,
				Updatedat: dd.Sizee.UpdatedAt,
			},
			SizeeID: dd.SizeeID,
			Createdat: dd.CreatedAt,
			Updatedat: dd.UpdatedAt,
			ColorID: dd.ColorID, 
			Quantity: dd.Quantity,
			Price: dd.Price,
			ProductImageID: dd.ProductImageID,
			Product_Images: domain.Product_Images{
				ID: dd.Product_Images.ID, 
				Image1: dd.Product_Images.Image1,
				Image2: dd.Product_Images.Image2,
				Image3: dd.Product_Images.Image3,
				Image4: dd.Product_Images.Image4,
			},
		})
	}
	return nil
}