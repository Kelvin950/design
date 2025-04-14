package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{

}



 func NewConfig()( *Config,error){

  err := godotenv.Load("../.env")

  if err!=nil{
	return nil , err
  }

  return &Config{} ,nil
 }


 func(c Config)GetEnv(key string)(string){

	return os.Getenv(key)
 }