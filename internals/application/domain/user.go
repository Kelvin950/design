package domain

import (
	"fmt"
	"time"
)

type User struct{
	ID uint  `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	UserType string `json:"user_type"`
	FirebaseID string `json:"firebase_id"`
	UserRole string `json:"user_role"`
	Avatar string	`json:"avatar"`
	Email string 	`json:"email"`
	Password string `json:"password"`
	Orders []Orders `json:"orders"`
	CreatedAt  time.Time 	`json:"createdat"`
	UpdatedAt time.Time 	`json:"updatedat"`
}


type FirebaseReturn struct {
    IDToken      string `json:"idToken"`
    Email        string `json:"email"`
    RefreshToken string `json:"refreshToken"`
    ExpiresIn    string `json:"expiresIn"`
    LocalID      string `json:"localId"`
    Registered   bool   `json:"registered"`
}

type UserLogin  struct{
	Email string `json:"email"`
	Password string `json:"password"`
	ReturnSecureToken bool 	`json:"returnSecureToken"`
}
type ApiError struct{
 
	Code int 
	Msg string	
}


func (e ApiError)Error()string{

	return  fmt.Sprintf("%s %d" , e.Msg , e.Code)
}