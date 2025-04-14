package db

import "gorm.io/gorm"

type OrderDetail struct{
	gorm.Model	
	ProductID *uint `gorm:"column:product_id"`
	Quantity int  
	Price float64
	OrderID *uint  `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}