package domain

import "time"

type Sizee struct {
	Name      string    `json:"name"`
	Id        uint      `json:"id"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}