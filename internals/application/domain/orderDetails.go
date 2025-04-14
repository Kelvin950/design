package domain

import "time"


type OrderDetail struct{
 Id        uint      `json:"id"`
ProductID *uint `json:"product_id"`
	Quantity int  `json:"quantity"`
	Price float64 	`json:"price"`
	OrderID *uint `json:"order_id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}