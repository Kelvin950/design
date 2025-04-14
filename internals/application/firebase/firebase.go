package firebase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	firebaseapp "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/kelvin950/desing/internals/application/domain"
)

type Firebase struct {
	App  *firebaseapp.App
	Key string
}

func Newfirebase(app *firebaseapp.App , key string)*Firebase{
	 

	return  &Firebase{
		App: app,
		Key: key,
	}
}

func (f Firebase) Createuser(user *domain.User) error {
 
   authclient , err:= f.App.Auth(context.Background())

   if err!=nil{
	return err
   }

   params := (&auth.UserToCreate{}).Email(user.Email).Password(user.Password)

  firebaseUser , err:= authclient.CreateUser(context.Background() ,params)
    
  if err!=nil{
	return err
  }
 
  user.FirebaseID =  firebaseUser.UID

  return nil

}

func (f Firebase)LoginUser(user *domain.User)(string , error){
	  


	var userlogin  = domain.UserLogin{
		Email: user.Email,
		Password: user.Password,
		ReturnSecureToken: true,
	}
	  p , err:= json.Marshal(userlogin)

	  if err!=nil{
		return "", err
	  }
 
	  buf:= bytes.NewBuffer(p)
     resp , err:= http.Post(fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s" ,f.Key) , "application/json" , buf)
		if err!=nil{
				return "" , err
			}
	
	  
	 if resp.StatusCode != http.StatusOK{
		
		return "" , domain.ApiError{
			Code: resp.StatusCode,
			Msg: "Incorrect Email or Password",
		}
	 }
 defer resp.Body.Close() 
     p ,err= io.ReadAll(resp.Body)
		if err!=nil{
				return "" , err
			}

			
	 var firebaseret domain.FirebaseReturn
	 err = json.Unmarshal(p , &firebaseret)

	   if err!=nil{
		return "" , err
	  }
	    authclient , err:= f.App.Auth(context.Background())
		if err!=nil{
				return "" , err
			}
		token , err := authclient.VerifyIDToken(context.TODO() , firebaseret.IDToken)
			 if err!=nil{
		return "" , domain.ApiError{
			Code: http.StatusBadRequest,
			Msg:  "Incorrect Email or Password",
		}
	  }
	return token.UID , nil
}