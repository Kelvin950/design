package db

import (
	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Typpe struct {
	gorm.Model 
	Name string 
	CategoryID  *uint  `gorm:"column:category_id"`
	AudienceID *uint 	`gorm:"column:audience_id"`
	UnitOfMeasurement *string  `gorm:"column:unit_of_measurement"`
	Products []Product  `gorm:"foreignKey:TyppeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


func(d DB)CreateTypes(t *domain.Typpe)error{

	newType := Typpe{
		Name: t.Name, 
		CategoryID: t.CategoryID, 
		AudienceID: t.AudienceID, 
		UnitOfMeasurement: t.UnitOfMeasurement, 
	}

	result := d.db.Save(&newType) 
	if result.Error !=nil{
		return result.Error 
	}
 
	t.Id = newType.ID
	t.Createdat = newType.CreatedAt 
	t.Updatedat = newType.UpdatedAt
	return nil 
}

