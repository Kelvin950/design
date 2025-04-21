package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{

}



 func NewConfig()( *Config){

  err := godotenv.Load("../.env")

fmt.Println(err,"=========================")

  return &Config{} 
 }


 func(c Config)GetEnv(key string)(string){

	return os.Getenv(key)
 }