package db

import (
	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Sizee struct {
	gorm.Model
	Name string
	ProductDetails []Product_Detail  `gorm:"foreignKey:SizeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	
}



func(d DB)CreateSizee(sizee *domain.Sizee)error{
	var size = Sizee{
		Name: sizee.Name,
	}
 
	 result:= d.db.Save(&size) 

	 if result.Error!=nil{
		return result.Error
	 }
 

	 sizee.Id = size.ID
	 sizee.Createdat =  size.CreatedAt
	 sizee.Updatedat = size.UpdatedAt
	 return nil 
}