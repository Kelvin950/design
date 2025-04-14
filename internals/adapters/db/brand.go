package db

import (
	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Brand struct {
	gorm.Model 
	Name string 
  Products  []Product `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


func(d DB)CreateBrand(brand *domain.Brand)error{
  var newBrand = Brand{
	Name:  brand.Name,
  }

  result := d.db.Save(&newBrand)
   
  if result.Error !=nil{
	 return result.Error 
  }

  brand.Id =  newBrand.ID 
  brand.Createdat =  newBrand.CreatedAt
  brand.Updatedat =  newBrand.UpdatedAt

  return  nil
}