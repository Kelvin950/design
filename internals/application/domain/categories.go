package domain

import "time"

type Category struct {
	Name      string    `json:"name"`
	Id        uint      `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
	Types   []Typpe  `json:"types"` 
}