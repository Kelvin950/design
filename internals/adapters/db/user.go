package db

import (
	"errors"
	"net/http"

	"github.com/kelvin950/desing/internals/application/domain"
	"gorm.io/gorm"
)

type Userr struct {
	gorm.Model 
	Email string 
	Fullname string
	UserType string  `gorm:"column:user_type"`
	UserRole string `gorm:"column:user_role"`
	Avatar string
	FirebaseID string `gorm:"column:firebase_id"`
	UserName string `gorm:"unique;column:username"`
	Orders  []Order
	Wishlist  []Wishlist  `gorm:"foreignKey:UserrID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (d DB)CreateUser(user *domain.User)error{

	newUser := &Userr{
		Email: user.Email,
		Fullname: user.Fullname,
		UserName: user.Username,
		UserRole: user.UserRole,
		UserType: user.UserType,
		FirebaseID: user.FirebaseID,
		Avatar: user.Avatar,
	}

	result := d.db.Save(newUser) 
	 
	if result.Error!=nil{
		return result.Error
	}

	user.ID = newUser.ID
	user.CreatedAt =  newUser.CreatedAt
	user.UpdatedAt =  newUser.UpdatedAt
	
	return nil
}

func(d  DB)GetUsers(users *[]domain.User)(error){
 
	
	var dd []Userr
	result := d.db.Find(&dd) 
	if result.Error!=nil{
		return   result.Error
	} 
 
	for _ , r:= range dd{
		
		*users =  append(*users , domain.User{
        ID: r.ID, 
		
		CreatedAt: r.CreatedAt, 
		UpdatedAt: r.UpdatedAt, 
		UserType: r.UserType,
			Fullname: r.Fullname,
		Email: r.Email,
		Avatar: r.Avatar,
		Username: r.UserName,
		})
	}

	return nil
}

func(d DB)GetUser(user *domain.User)(error){
	
	var dd Userr
	result := d.db.Find(&dd, user.ID) 
	if result.Error !=nil{
		return result.Error
	}
	user.Username= dd.UserName
	user.Avatar= dd.Avatar
	user.Fullname= dd.Fullname
	user.UserType=dd.UserType

	user.Email =  dd.Email
	user.CreatedAt= dd.CreatedAt
	user.UpdatedAt=  dd.UpdatedAt
	return nil
}

func(d DB)GetUserBYEmail(user *domain.User)error{
	var dd Userr
	result := d.db.First(&dd ,Userr{Email: user.Email}) 
	if result.Error !=nil{
		 if errors.Is(result.Error , gorm.ErrRecordNotFound){
			return domain.ApiError{Code: http.StatusNotFound , Msg: "user not found"}
		 }

		 return result.Error
	}
	user.Username= dd.UserName
	user.Avatar= dd.Avatar
	
	user.Email =  dd.Email
	user.CreatedAt= dd.CreatedAt
	user.UpdatedAt=  dd.UpdatedAt
	return nil
}

func(d DB)GetUserByUserName(user *domain.User)error{
		var dd Userr
	result := d.db.First(&dd ,Userr{Email: user.Username}) 
	if result.Error !=nil{
		 if errors.Is(result.Error , gorm.ErrRecordNotFound){
			return  domain.ApiError{Code: http.StatusNotFound ,Msg: "user not found"}
		 }

		 return result.Error
	}
	
	
	user= &domain.User{
		Username: dd.UserName, 
		Avatar: dd.Avatar,
		Email: dd.Email,
		CreatedAt: dd.CreatedAt,
		UpdatedAt: dd.UpdatedAt,
		
	}
	return nil
}



func(d DB)GetUserByFirebaseUID(user *domain.User)error{
		var dd Userr
	result := d.db.First(&dd ,Userr{FirebaseID: user.FirebaseID}) 
	if result.Error !=nil{
		 if errors.Is(result.Error , gorm.ErrRecordNotFound){
			return  domain.ApiError{Code: http.StatusNotFound ,Msg: "user not found"}
		 }

		 return result.Error
	}
	

	user.Username= dd.UserName
	user.Avatar= dd.Avatar
	user.ID = dd.ID
	user.Password="" 
	user.UserRole= dd.UserRole
	user.UserType= dd.UserType
	
	user.Email =  dd.Email
	user.CreatedAt= dd.CreatedAt
	user.UpdatedAt=  dd.UpdatedAt
	
	
	return nil
}