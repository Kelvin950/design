package api

import "github.com/kelvin950/desing/internals/application/domain"


func(a Api)GetUserWishlist(userid uint)([]domain.Wishlist  , error){
 

	var wishlist =  []domain.Wishlist{
		{
			UserID: &userid,
		} ,
	}
 

	userwishlist ,err:= a.DB.GetWishlistByUserID(wishlist)

	 if err!=nil{

		return nil , domain.ApiError{Code:500 , Msg:err.Error()}
	 }


	 return userwishlist , nil
}




// func(a Api)CreateUser(user *domain.User)error{


// 	err:= a.DB.CreateUser(user) 
	
// 	 if err!=nil{
// 		return err
// 	 }

// 	 return nil 
// } 