package api

import (
	
	"os"
	"testing"

	"github.com/kelvin950/desing/internals/application/domain"
	"github.com/stretchr/testify/require"
)


type test struct{
	t Token
	user domain.User
}

 var test1 = test{
	t:  Token{
	secret: []byte("1232"),
	
} ,

 }




func TestDecode(t *testing.T) {
	docodedToken, err:= test1.t.Decode("eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ5NzQwYTcwYjA5NzJkY2NmNzVmYTg4YmM1MjliZDE2YTMwNTczYmQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI3MzgyNjYxMTU5NDQtc2FnOWRzazhlYm9paGI3ZnBlNmNwMG5zcGFxZWp1YzIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI3MzgyNjYxMTU5NDQtc2FnOWRzazhlYm9paGI3ZnBlNmNwMG5zcGFxZWp1YzIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDkxNjI3NjM1NjQ0NzI5ODcxMjQiLCJlbWFpbCI6ImRlbmxpbmF0b0BnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmJmIjoxNzMyNDYyMzQ1LCJuYW1lIjoiS2VsdmluIEF0byBCYWlkZW4iLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jTFA0NkZWTzRselRTbGFmejY4SkFnVnhNdTNHcGtUNjVXeWp0WFp6eHBIbW1Fa0N5NTc9czk2LWMiLCJnaXZlbl9uYW1lIjoiS2VsdmluIEF0byIsImZhbWlseV9uYW1lIjoiQmFpZGVuIiwiaWF0IjoxNzMyNDYyNjQ1LCJleHAiOjE3MzI0NjYyNDUsImp0aSI6IjJjMThmODRhNmUwYWNiNjNmODY1MmZmYWU1MTg4OTMwNWVjZmU2YjkifQ.iWC1YL6FLS2oyaxYqy72Se5_64_lN-v4TwZCtEDjHvQSvCUuiUoW0tIn9irBu5vhz20bvYpPV4GH6cdPalENLK1AApcGA-QAEZL2s_30wF-pSwM2R1w6jOAG9ICLFN36AlX91acgKGdOiygWm7Sz4vUPLShLIi7EuaPZQn1gYK6C0Ww-72HW-z5tvMMTERJyrSCbhhqRo6jPrN3a27q38B7YFluXZQ1-Mb4bUnNC1MicgvL5jQane4LZf1Ht1UYL5r82LMVDaDUIBrQTUZGLRgtqjqHthgNYnOSk9fGY2XfPc2tHTv0VbufLa1HDrrUeFjaUIqbem1aDJoMUihC_vQ")
	require.NoError(t , err) 
	require.NotEmpty(t , docodedToken)  
	t.Log(docodedToken.Email , docodedToken.EmailVerified , docodedToken.FamilyName , docodedToken.GivenName, docodedToken.Picture)
	test1.user =  domain.User{
		Email: docodedToken.Email, 
	    Username: "dsds", 
		Fullname: docodedToken.GivenName,
	
		Avatar: docodedToken.Picture,

	}
}

func TestSign(t *testing.T){
	 
	token ,err:= test1.t.Sign(test1.user)
	require.NoError(t ,err) 
	require.NotEmpty(t , token)

 user ,err:=	test1.t.Verify(token)
t.Log(token)
 require.NoError(t ,err)
  require.Equal(t ,user.Email ,test1.user.Email)
}




func TestMain(m *testing.M){


	v :=m.Run()
	os.Exit(v)
}


