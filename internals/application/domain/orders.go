package domain

import "time"

type Orders struct {
	Id            uint      `json:"id"`
	Createdat     time.Time `json:"created_at"`
	Updatedat     time.Time `json:"updated_at"`
	UserrID       *uint     `json:"user_id"`
	TotalQuantity int       `json:"total_quantity"`
	TotalPrice    float64   `json:"total_price"`
	OrderDetails []OrderDetail `json:"order_details"`
}