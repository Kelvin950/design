package db

import "gorm.io/gorm"

type OrderDetail struct{
	gorm.Model	
	ProductID *uint `gorm:"column:product_id"`
	Product Product`gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity int  
	Price float64
	OrderID *uint  `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}