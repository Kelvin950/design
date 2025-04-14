package domain

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"os"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestH(t *testing.T){
	var x bytes.Buffer 
	
	h:= H{}
	h.Add("fd" , "dsc")
	
	
	n ,err:= h.WriteTo(os.Stdout)
	require.NoError(t, err) 
	
	t.Log(n)
	t.Logf("%s" , x.Bytes())
  
}



func TestQ(t *testing.T){

	req, err:=http.NewRequest("POST", "http://localhost:3001/loginUser" ,strings.NewReader("{credentials: eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ5NzQwYTcwYjA5NzJkY2NmNzVmYTg4YmM1MjliZDE2YTMwNTczYmQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI3MzgyNjYxMTU5NDQtc2FnOWRzazhlYm9paGI3ZnBlNmNwMG5zcGFxZWp1YzIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI3MzgyNjYxMTU5NDQtc2FnOWRzazhlYm9paGI3ZnBlNmNwMG5zcGFxZWp1YzIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDkxNjI3NjM1NjQ0NzI5ODcxMjQiLCJlbWFpbCI6ImRlbmxpbmF0b0BnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmJmIjoxNzMyNDYyMzQ1LCJuYW1lIjoiS2VsdmluIEF0byBCYWlkZW4iLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jTFA0NkZWTzRselRTbGFmejY4SkFnVnhNdTNHcGtUNjVXeWp0WFp6eHBIbW1Fa0N5NTc9czk2LWMiLCJnaXZlbl9uYW1lIjoiS2VsdmluIEF0byIsImZhbWlseV9uYW1lIjoiQmFpZGVuIiwiaWF0IjoxNzMyNDYyNjQ1LCJleHAiOjE3MzI0NjYyNDUsImp0aSI6IjJjMThmODRhNmUwYWNiNjNmODY1MmZmYWU1MTg4OTMwNWVjZmU2YjkifQ.iWC1YL6FLS2oyaxYqy72Se5_64_lN-v4TwZCtEDjHvQSvCUuiUoW0tIn9irBu5vhz20bvYpPV4GH6cdPalENLK1AApcGA-QAEZL2s_30wF-pSwM2R1w6jOAG9ICLFN36AlX91acgKGdOiygWm7Sz4vUPLShLIi7EuaPZQn1gYK6C0Ww-72HW-z5tvMMTERJyrSCbhhqRo6jPrN3a27q38B7YFluXZQ1-Mb4bUnNC1MicgvL5jQane4LZf1Ht1UYL5r82LMVDaDUIBrQTUZGLRgtqjqHthgNYnOSk9fGY2XfPc2tHTv0VbufLa1HDrrUeFjaUIqbem1aDJoMUihC_vQ}"))

    require.NoError(t, err)
	res, err :=http.DefaultClient.Do(req)

    require.NoError(t, err)
	defer res.Body.Close() 
	 
	 p,err:=  io.ReadAll(res.Body)
	 require.NoError(t, err)
	t.Log(p)
}


