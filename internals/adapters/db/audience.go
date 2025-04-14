package db

import (
	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Audience struct {
	gorm.Model 
	Name string
	Types []Typpe
}

func(a DB)CreateAudience(audience  *domain.Audience)error{

   newAudience := &Audience{
	    Name: audience.Name  ,
   }
    result := a.db.Save(newAudience) 
   if result.Error != nil{
	return result.Error
   }

  audience.Createdat=  newAudience.CreatedAt 
  audience.Updatedat =  newAudience.UpdatedAt 
  audience.Id =  newAudience.ID 
  return  nil 
}


