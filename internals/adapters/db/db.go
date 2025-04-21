package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DB struct {
	db *gorm.DB
}

func NewDb(dsn string) (*DB ,error) {
 
 

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		  NamingStrategy: schema.NamingStrategy{
        SingularTable: true, // Use singular table names
    },
	} , )

	if err !=nil{
		return nil , err
	}

	err =  db.AutoMigrate(&Userr{}  ,&Audience{} , &Category{} ,&Typpe{}  , &Brand{} , &Sizee{} , &Color{} , &Product{} , &Product_Detail{} , &Order{} , &OrderDetail{} , &Wishlist{}) 
   db = db.Debug()
	if err !=nil{
		return nil , err
	}

	return &DB{
		db: db,
	} ,nil
}