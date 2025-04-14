package domain

import "time"

type Typpe struct {
	Id uint `json:"id"` 
	Createdat time.Time   `json:"created_at"`
	Updatedat time.Time  `json:"updated_at"`
	Name string `json:"name"`
	CategoryID *uint `json:"category_id"`
	AudienceID *uint  `json:"audience_id"`
	UnitOfMeasurement *string  `json:"unit_of_measurement"`
}