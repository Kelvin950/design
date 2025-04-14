package domain

import (
	"encoding/binary"
	"encoding/json"
	"io"
)


type H map[string]interface{} 


func(h H)Add(key string , val interface{}){
  
	 h[key] =  val
}

func (h H)WriteTo( w io.Writer)(int64,error){
   
	 p,err:=json.Marshal(h)
	 if err!=nil{
		return 0 , err
	 }

	 size := uint32(len(p)) 

	//  err= binary.Write(w , binary.BigEndian , size)
	// if err!=nil{
	// 	return 0 , err
	// }

	  n,err:= w.Write(p)

	  if err!=nil{
		return 0 , err
	}

	 return int64(binary.Size(size) + n ) , nil
}