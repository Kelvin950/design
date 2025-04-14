package domain

import "time"

type Color struct {
	Name      string    `json:"name"`
	Code string `json:"code"`
	Id        uint      `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}