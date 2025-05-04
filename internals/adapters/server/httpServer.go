package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/kelvin950/desing/internals/application/domain"
	"github.com/kelvin950/desing/internals/ports"
)

type Server struct {
	Router *mux.Router
	Api ports.Api
}



func NewServer(api ports.Api)*Server{


	return  &Server{
		Api: api,
	}
}


func (s *Server)Routes(){
 
	router := mux.NewRouter() 

	router.Handle("/loginUser", s.SiginUser()).Methods("POST")
	router.Handle("/createUser",s.CreateUser() ).Methods("POST")
	router.Handle("/products", s.CreateProduct()).Methods("POST")
	router.Handle("/products", s.GetProduct()).Methods("GET")
	router.Handle("/products_detail" , s.CreateProduct_Detail()).Methods("POST")
	router.Handle("/products_detail", s.GetProductDetails()).Methods("GET")
	router.Handle("/order" , s.CreateOrders()).Methods("POST")
	// router.Handle("/order" , s.GetOrders()).Methods("GET")
	router.Handle("/order" , s.GetUserOrder()).Methods("GET")
	router.Handle("/wishlist" , s.GetUserWishlist()).Methods("GET")
	router.Handle("/wishlist"  , s.CreateWishlist()).Methods("POST")
  s.Router =  router
}


func (s *Server)CreateUser()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		var  req domain.User 
		if err:= json.NewDecoder(r.Body).Decode(&req); err!=nil{
			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return 
		}

		err:= s.Api.CreateUser(&req)  

		if err!=nil{
			w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
		}
  

		w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)
			res := domain.H{
			 "user":req,
			}
			res.WriteTo(w)
		

	}
}

func(s *Server)SiginUser()http.HandlerFunc{ 
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		var req domain.User 
		 if err:= json.NewDecoder(r.Body).Decode(&req);  err!=nil{
			w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return 
	}


	  err:= s.Api.SigInUser(&req)
  if err!=nil{

	  if apiError , ok := err.(domain.ApiError) ; ok{
		w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(apiError.Code)
			res := domain.H{
			 "message": apiError.Msg,
			}
			res.WriteTo(w)
			return 
	  }
		w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message":err.Error(),
			}
			res.WriteTo(w)
			return 
	  
  }
	  
	w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)
			res := domain.H{
			 "user":req,
			}
			res.WriteTo(w)
		

}

}

func(s *Server)LoginUser()http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		var req domain.LoginRequest
		 err:= json.NewDecoder(r.Body).Decode(&req); if err!=nil{

			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return 
		 }

		 user ,token, err:=s.Api.LoginUser(req.Credentials , "738266115944-sag9dsk8eboihb7fpe6cp0nspaqejuc2.apps.googleusercontent.com" )
		 if err!=nil{
			apiError ,ok:= err.(domain.ApiError) 
			log.Println(err)
			if !ok{
				
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
 
			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(apiError.Code) 
			 res := domain.H{
				"success":false ,
				 "message":apiError.Msg ,
			 }
			 res.WriteTo(w)
			return 
		 }
		 

		  w.Header().Add("Content-Type" ,"application/json")
		 w.WriteHeader(http.StatusOK) 
		 res := domain.H{
			"success":true ,
			"data":user ,
			  "token":token ,
		 }
		res.WriteTo(w) 
		
	}
}


func  (s Server)CreateProduct()http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		 var req domain.Product 
		
		  err:= json.NewDecoder(r.Body).Decode(&req); if err!=nil{
			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
		  }

		
	err = s.Api.CreateProduct(&req)

	if err!=nil{
		w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
	}

    w.Header().Add("Content-Type" ,"application/json")
		 w.WriteHeader(http.StatusOK) 
		 res := domain.H{
			"success":true ,
			 "data":req ,
		 }
		 res.WriteTo(w)
	}
}


func (s Server)GetProduct()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
 
		var products  []domain.Product 

		err:= s.Api.GetProducts(&products) 
		if err!=nil{
				w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			fmt.Println(err)
			res.WriteTo(w)
			return
		}

		    w.Header().Add("Content-Type" ,"application/json")
		 w.WriteHeader(http.StatusOK) 
		 res := domain.H{
			"success":true ,
			"products" : products ,
		 }
		 res.WriteTo(w)

	}
}

func(s Server)CreateProduct_Detail()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) { 
		var req domain.Product_Detail
		 if err:= json.NewDecoder(r.Body).Decode(&req);err!=nil{
			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": err,
			}
			fmt.Println(err)
			res.WriteTo(w)
			return
		 }

		 
		 err:= s.Api.CreateProduct_Detail(&req) 
		 if err!=nil{
				 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
		 }

	  	 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)
			res := domain.H{
			 "product_detail": req,
			}
			res.WriteTo(w)
	}
}


func(s Server)GetProductDetails()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		var products  []domain.Product_Detail

	prod_dets ,	err:= s.Api.GetProductDetails(&products) 
		if err!=nil{
				w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
		}

		    w.Header().Add("Content-Type" ,"application/json")
		 w.WriteHeader(http.StatusOK) 
		 res := domain.H{
			"success":true ,
			"products" : products ,
			"dets" : prod_dets ,
		 }
		 res.WriteTo(w)

	}
}


func(s Server)CreateOrders()http.HandlerFunc{
	
	return func(w http.ResponseWriter, r *http.Request) {

		var req  []domain.OrderDetail

		if err:= json.NewDecoder(r.Body).Decode(&req) ;err!=nil{
			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
		}
	
		 
		userId:= r.URL.Query().Get("userId")

       i,err:= strconv.Atoi(userId)
	  if err!=nil{
			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": "user id incorrect",
			}
			res.WriteTo(w)
			return
		}
	
		orders , err:= s.Api.CreateOrders(&req , i )

		if err!=nil{
				 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
		}


		 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)
			res := domain.H{
			  "orders": orders,
			}
			res.WriteTo(w)
		
 	}
}


func(s Server)GetOrders()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

 
		var orders []domain.Orders
		err:= s.Api.GetOrders(&orders) 

		if err!=nil{
				 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			res.WriteTo(w)
			return
		}



		 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)
			res := domain.H{
			  "orders": orders,
			}
			res.WriteTo(w)

	}
}


func(s Server)GetUserOrder()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		
		
		userId := r.URL.Query().Get("userId")
		 
		i , err:=strconv.Atoi(userId)

		if err!=nil{
				 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			fmt.Println(err)
			res.WriteTo(w)
			return
		}
 
       userId1 := uint(i)
		var order =  domain.Orders{
			UserrID: &userId1,
		}
		err = s.Api.GetUserOrder(&order)
 	if err!=nil{

		if apiError , ok := err.(domain.ApiError) ; ok{
			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(apiError.Code)
			res := domain.H{
			 "message": apiError.Msg,
			}
			fmt.Println(err)
			res.WriteTo(w)
			return
		}
		
				 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			fmt.Println(err)
			res.WriteTo(w)
			return
		}
 
w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)
			res := domain.H{
			  "orders": order,
			}
			res.WriteTo(w)


	}


	
}

func(s Server)CreateWishlist()http.HandlerFunc{
	return   func(w http.ResponseWriter, r *http.Request) {
		
		var req domain.Wishlist ; 

		if err := json.NewDecoder(r.Body).Decode(&req); err !=nil{

			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusBadRequest)
			res := domain.H{
			 "message": err,
			}
			
			res.WriteTo(w)
			return
		} 
  
	 
		dd:=&[]domain.Wishlist{req}

		err:= s.Api.CreateWishlist(dd ) 
		
		if err!=nil{


			 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			
			res.WriteTo(w)
			return
		}

	 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)		 
		res:= domain.H{
			"success":true ,
			"wishlist":dd ,
		}
		res.WriteTo(w)

	}
}

func(s Server)GetUserWishlist()http.HandlerFunc{

     return func(w http.ResponseWriter, r *http.Request) {

		userId := r.URL.Query().Get("userId")
	  
		i , err:=strconv.Atoi(userId)

		if err!=nil{
				 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			fmt.Println(err)
			res.WriteTo(w)
			return
		}
 
	
		userWishlist , err := s.Api.GetUserWishlist(uint(i)) 
 
		if err!=nil{
				 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusInternalServerError)
			res := domain.H{
			 "message": err,
			}
			fmt.Println(err)
			res.WriteTo(w)
			return
		}
 
		 w.Header().Add("Content-Type" ,"application/json")
			w.WriteHeader(http.StatusOK)		 
		res:= domain.H{
			"success":true ,
			"wishlist":userWishlist ,
		}
		res.WriteTo(w)
	 }
	

}