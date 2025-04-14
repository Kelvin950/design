package api

import (
	"net/http"


	"github.com/kelvin950/desing/internals/application/domain"
	"github.com/kelvin950/desing/internals/ports"
)

type Api struct {
	DB ports.DB
	*Token
	Firebase  ports.Firebase
}

func  NewApi(secret string , db ports.DB , firebase ports.Firebase)*Api{

	jwt := NewToken(secret) 

	return &Api{
		
		DB: db ,
		Token: jwt,
		Firebase: firebase,
	}

}

func (a Api)LoginUser(cred string , aud string)( domain.User,string , error){ 
  
	claims ,err:= a.Decode(cred)
	if err!=nil{
		return domain.User{},"", domain.ApiError{Code:http.StatusBadRequest , Msg:err.Error()}
	} 
	if claims.Aud != aud{
		return domain.User{ } ,"", domain.ApiError{Code:http.StatusBadRequest , Msg:"cred invalid"}
	}
   
	//find user by email 
	 user := domain.User{
		Email: claims.Email,
	 }

	err = a.DB.GetUserBYEmail(&user) 
	 
	if err!=nil{
			_ ,ok:= err.(domain.ApiError) 
	if !ok{
		return  domain.User{ } ,"", err
	}
 
		user= domain.User{
		Email: claims.Email, 
		Fullname: claims.GivenName,
		
		Avatar: claims.Picture,
	  }
	 err = a.DB.CreateUser(&user)

	  if err!=nil{
		return  domain.User{ } ,"" ,err
	  }
 
	 token , err :=a.Token.Sign(user)
  	if err!=nil{
		return  domain.User{ } , "" ,err
	  }

	   return user , token , nil
	}
	
	 token , err :=a.Token.Sign(user)
  	if err!=nil{
		return  domain.User{ } , "" ,err
	  }
	  return user , token , nil
 
}

func(a  *Api)CreateUser(user *domain.User)error{
  
	err:=a.Firebase.Createuser(user)

	if err!=nil{
		return err
	}


	err = a.DB.CreateUser(user) 
	if err!=nil{
		return err
	}

	return nil
}

func(a Api)SigInUser(user *domain.User)(error){
	
   firebaseUid , err:= a.Firebase.LoginUser(user)

   if err!=nil{
	return err
   }


   user.FirebaseID = firebaseUid 
 
  err = a.DB.GetUserByFirebaseUID(user)
    if err!=nil{
	return err
   }

	 return nil
}