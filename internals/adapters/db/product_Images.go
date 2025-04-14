package db

import "gorm.io/gorm"

type Product_Images struct {
	gorm.Model 
	Image1 string  `gorm:"column:image_1"`
	Image2 string 	`gorm:"column:image_2"`
	Image3 string  `gorm:"column:image_3"`
	Image4 string  `gorm:"column:image_4"`
	Product_Details []Product_Detail `gorm:"foreignKey:ProductImageID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}