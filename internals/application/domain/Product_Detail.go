package domain

import "time"

type Product_Detail struct {
	Id       uint    `json:"_id"`
	ProductID *uint      `json:"product_id"`
	ColorID   *uint      `json:"color_id"`
	SizeeID   *uint      `json:"size_id"`
	Product  Product 	  `json:"product"`
	Color 	Color 		 `json:"color"` 
	Sizee	Sizee 		`json:"size"`
	Quantity  int      `json:"quantity"`
	Price    float64   `json:"price"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	ProductImageID *uint      `json:"product_image_id"`
	Product_Images Product_Images `json:"product_images"`
}


type Product_Images struct{
	ID uint `json:"_id"` 
	Image1 string  `json:"column:image_1"`
	Image2 string 	`json:"column:image_2"`
	Image3 string  `json:"column:image_3"`
    Image4 string  `json:"column:image_4"`
}


type Product_Detail_Colors struct{
	Name string  `json:"name"` 
	Code string `json:"code"`
	ID uint `json:"_id"` 
	Sizes  []Product_Detail_Size  `json:"sizes"`
}

type Product_Detail_Product  struct{
	Name string  `json:"name"`
	ID uint 	`json:"_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Quantity  int `json:"quantity"` 
	Price  float64  `json:"price"` 
	Colors []Product_Detail_Colors `json:"colors"` 
	ImageUrl  []string `json:"image_url"`
}

type Product_Detail_Size struct{
	Name string  `json:"name"`
	Quantity int `json:"quantity"`
	Price float64  `json:"price"`
}