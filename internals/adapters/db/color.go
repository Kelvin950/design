package db

import (
	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Color struct {
	gorm.Model 
	Name string 
	Code string 
	ProductDetails []Product_Detail  `gorm:"foreignKey:ColorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


func(d DB)CreateColor(color *domain.Color)error{
 
	var newColor  = Color{
		Name:  color.Name, 
		Code: color.Code,
	}

	result:= d.db.Save(&newColor) 
	
	if result.Error!=nil{
		return  result.Error
	}

	color.Id =  newColor.ID 
	color.Createdat =  newColor.CreatedAt 
	color.Updatedat =  newColor.UpdatedAt
	return nil 
}
