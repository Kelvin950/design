package db

import (
	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model 
	Name string 
   Types   []Typpe
}


func(d DB)CreateCategory(category *domain.Category)error{
	 
  newCategory := Category {
	Name: category.Name,
  }

  result :=d.db.Save(&newCategory) 

  if result.Error!= nil{
	return result.Error 
  }


  category.Id =  newCategory.ID 
  category.Updatedat =  newCategory.UpdatedAt
  category.Createdat = newCategory.CreatedAt 

  return nil 

}

