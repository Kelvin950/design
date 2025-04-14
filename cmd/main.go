package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	

	"time"

	firebaseapp "firebase.google.com/go/v4"
	"github.com/kelvin950/desing/config"
	"github.com/kelvin950/desing/internals/adapters/db"
	"github.com/kelvin950/desing/internals/adapters/server"
	"github.com/kelvin950/desing/internals/application/api"
	"github.com/kelvin950/desing/internals/application/firebase"
	"golang.org/x/sys/unix"
	"google.golang.org/api/option"
)

func main() {

 
	configr,err  := config.NewConfig()
	if err!=nil{
		log.Fatal(err) 
	}

	 dbUser := configr.GetEnv("DB_USER")
	 dbpassword:= configr.GetEnv("DB_PASSWORD")
	 DBName :=  configr.GetEnv("DB_Name")
	 dbPort :=  configr.GetEnv("DB_PORT")
	firebasekey:= configr.GetEnv("FIREBASE_KEY")
	dsn :=fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai" , dbUser , dbpassword , DBName ,dbPort)
 dbConn, err :=db.NewDb(dsn)
 if err!=nil{
	log.Fatal(err)
 } 
 
opt := option.WithCredentialsFile("../design-ff9af-firebase-adminsdk-fbsvc-7decff48f7.json")
app, err := firebaseapp.NewApp(context.Background(), nil, opt)
if err != nil {
	log.Fatalf("error initializing app: %v", err)
}


  firebasev := firebase.Newfirebase(app  , firebasekey)
 
   api := api.NewApi("12323" ,dbConn ,firebasev) 
   
   httpServer := server.NewServer(api) 
   httpServer.Routes() 
  
   serve:= http.Server{
	Addr: ":3001", 
	Handler: httpServer.Router, 
	
   } 
 
   done := make(chan error)
   
   go func(){

	err:= serve.ListenAndServe()
	if err!=nil && err!= http.ErrServerClosed{
       done <- err
	}
	close(done)
   }()

   

      sig:= make(chan os.Signal ,1) 
	  signal.Notify(sig , unix.SIGTERM, unix.SIGINT) 
	
	//    ctx , cancel:=  context.WithTimeout(context.Background() , 10* time.Second) 
	//  defer cancel()
	//   log.Println("server shutting down")
  
    //  serve.Shutdown(ctx)
   

	 select{
	 case   <-sig:
		   ctx , cancel:=  context.WithTimeout(context.Background() , 10* time.Second) 
	 defer cancel()
	  log.Println("server shutting down")
  
     serve.Shutdown(ctx)

	 case err:=<-done :
		log.Fatalf("server error %s" , err) 
		
	 
	 }
    

     
	

}