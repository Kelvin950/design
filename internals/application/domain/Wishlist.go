package domain

import "time"

type Wishlist struct {
	ID        uint      `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	UserID    *uint     `json:"user_id"`
	ProductId  uint `json:"product_id"`
	Products  Product `json:"products" `
}

