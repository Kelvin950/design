package api

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kelvin950/desing/internals/application/domain"
)

type CustomClaims  struct{
domain.User 
jwt.RegisteredClaims
}


type GoogleClaims struct{
	
	
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Aud           string `json:"aud"`
	jwt.RegisteredClaims
}
type Token struct{
 secret []byte  
}


func NewToken(jwtSecret string)*Token{
  
	return &Token{
		secret: []byte(jwtSecret), 
	}
}

func(t Token)Sign(user domain.User)(string ,error){

claims:= CustomClaims{
	user , 
	jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		Issuer: "design",
		IssuedAt: jwt.NewNumericDate(time.Now()),
	},
}

  token := jwt.NewWithClaims(jwt.SigningMethodHS256 , claims ) 
    return  token.SignedString(t.secret)

}

func(tq Token)Verify(token string)(domain.User , error){
  
	t , err:=jwt.ParseWithClaims(token , &CustomClaims{} , func(t *jwt.Token) (interface{}, error) {
		 return tq.secret, nil
	}) 
	if err!=nil{
		return  domain.User{}, err
	}

	
	claims , ok := t.Claims.(*CustomClaims)
	if !ok{
		return domain.User{}, errors.New("invalid token")
	}


	return claims.User , nil

}

func(t Token)Decode(token string)(GoogleClaims ,error){

	 parser :=jwt.NewParser() 
	v , _ , err:=parser.ParseUnverified(token , &GoogleClaims{})

	if err!=nil{
		return GoogleClaims{} , err
	}
    
	 claims ,ok:=v.Claims.(*GoogleClaims)
	 if !ok{
		return GoogleClaims{},errors.New("failed")	 }
	return  *claims ,nil
}