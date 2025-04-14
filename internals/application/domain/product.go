package domain

import "time"



type Product struct{
	 Id        uint      `json:"id"`
   Name string `json:"name"`
   Color []int64 `json:"color"` 
   Sizes []int64 `json:"sizes"` 
   Quantity int `json:"quantity"`
   MinPrice float64 `json:"min_price,omitempty"`
   MaxPrice float64  `json:"max_price,omitempty"`
   BrandID *uint `json:"brand_id,omitempty"` 
   Brand  Brand  `json:"brand,omitempty"`
   Typpe Typpe    `json:"type,omitempty"`
   TyppeID *uint `json:"type_id,omitempty"`
   OrderDetails []OrderDetail  `json:"order_details,omitempty"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}