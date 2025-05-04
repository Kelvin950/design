package db

import (
	"context"
	"fmt"
	"log"

	"strings"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/kelvin950/desing/internals/application/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)


type dbsuite struct{

	suite.Suite 
	*DB
container  testcontainers.Container
}


func (suite *dbsuite)SetupSuite(){
	port:="5432"
	req := testcontainers.ContainerRequest{
		Image: "postgres",
		ExposedPorts: []string{port}, 
		Env: map[string]string{
			"POSTGRES_PASSWORD":"gorm",
				"POSTGRES_USER":"gorm",
				"POSTGRES_DB":"design",
		},
		WaitingFor: wait.ForSQL(nat.Port(port), "pgx", func(host string, port nat.Port) string {
			return fmt.Sprintf("host=localhost user=gorm password=gorm dbname=design port=%s sslmode=disable TimeZone=Asia/Shanghai" , port.Port())
		}).
        WithStartupTimeout(time.Second * 20).
        WithQuery("SELECT 10"),
	} 

	container , err:= testcontainers.GenericContainer(context.Background() , testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started: true,
	}) 

	if err!=nil{
		log.Fatal(err)
	}

	suite.container =  container 
	endpoint,err:= container.Endpoint(context.Background() , "") 
	if err!=nil{
		log.Fatal(err)
	}
	db , err:= NewDb(fmt.Sprintf("host=localhost user=gorm password=gorm dbname=design port=%s sslmode=disable TimeZone=Asia/Shanghai" , strings.Split(endpoint, ":")[1]))
	if err!=nil{
		log.Fatal(err)
	}
  suite.DB =  db
}


func(suite *dbsuite)TearDownSuite(){
 

	suite.container.Terminate(context.Background())
    
}

func(suite *dbsuite)TestCreateUser(){
 var user  = &domain.User{
	Email: "fdfdfd", 
	Fullname: "ds",
	Avatar: "dsewe",
	Username: "_dee,w",
 }
	err:= suite.CreateUser(user) 

	suite.NoError(err)
 suite.NotNil(user) 
}


func(suite *dbsuite)TestGetUsers(){
	var user = &[]domain.User{}
	 
	err:= suite.GetUsers(user) 
	suite.Suite.NoError(err) 
	suite.Suite.NotEqual(len(*user), 0)
	suite.T().Log(user)
}

func(suite *dbsuite)TestGetUser(){
	var user =  &domain.User{ID: 1} 
	err:= suite.GetUser(user) 
	suite.NoError(err)
	suite.NotEmpty(user.Fullname )
}


func(suite *dbsuite)TestCreateAudience(){
	var audience=  domain.Audience{
		Name: "Women", 
		
	}
 

	
  err:= suite.CreateAudience(&audience) 
 suite.NotEmpty(audience.Createdat )
 suite.NotEmpty(audience.Updatedat) 
 suite.NotEmpty(audience.Id)
  suite.NoError(err) 
}

func(suite *dbsuite)TestCreateCategory(){

	var category =  domain.Category{
		Name: "Clothing",
	}
	err:=suite.CreateCategory(&category) 

	suite.T().Log("kll")
	suite.Greater(category.Id, uint(0))
	 suite.NotEmpty(category.Createdat) 
	 suite.NotEmpty(category.Updatedat)
	 suite.NoError(err)
}


func (suite  *dbsuite)TestTyppe(){
	var   category =  domain.Category{
		Name: "Clothing",
	}
	err:=suite.CreateCategory(&category) 

suite.NoError(err)
	var audience=  domain.Audience{
		Name: "Women", 
		
	}
 

	
  err = suite.CreateAudience(&audience) 

suite.NoError(err)
var m = "xs|s|M|L"
  var typpe =  domain.Typpe{
 		Name: "T-shirt", 
		CategoryID: &category.Id,
		AudienceID: &audience.Id, 
		UnitOfMeasurement: &m,
  } 

    err = suite.CreateTypes(&typpe) 

	suite.NoError(err)
	suite.NotEmpty(typpe.Id) 
	suite.Equal(typpe.AudienceID , &audience.Id) 
	suite.Equal(typpe.CategoryID , &category.Id)


}

func(suite *dbsuite)TestCreateProduct(){
		var   category =  domain.Category{
		Name: "Clothing",
	}
	err:=suite.CreateCategory(&category) 

suite.NoError(err)
	var audience=  domain.Audience{
		Name: "Women", 
		
	}
 

	
  err = suite.CreateAudience(&audience) 

suite.NoError(err)
var n = "xs|s|M|L"
  var typpe =  domain.Typpe{
 		Name: "T-shirt", 
		CategoryID: &category.Id,
		AudienceID: &audience.Id, 
		UnitOfMeasurement:&n ,
  } 

    err = suite.CreateTypes(&typpe) 

	suite.NoError(err)
	suite.NotEmpty(typpe.Id) 
	suite.Equal(typpe.AudienceID , &audience.Id) 
	suite.Equal(typpe.CategoryID , &category.Id)
 

	var brand=  domain.Brand{
		Name: "Adidas",
	}
	err=  suite.CreateBrand(&brand)
	suite.NoError(err) 
   
	var color = domain.Color{
		Name: "White",
		Code: "#fff",
	}

	err =  suite.CreateColor(&color) 
	suite.NoError(err) 

	size := domain.Sizee{
		Name: "XL",
	}

	err = suite.CreateSizee(&size) 
	suite.NoError(err) 

	product:= domain.Product{
		Name: "T-shirt",
		Color: []int64{int64(color.Id)},
		Sizes: []int64{int64(size.Id)} ,
		Quantity: 10, 
		MaxPrice: 100,
		MinPrice: 12, 
		BrandID: &brand.Id, 
		TyppeID: &typpe.Id,
	}

	err= suite.CreateProduct(&product) 
	suite.NoError(err) 

	suite.NotEmpty(product.Id) 
	suite.NotEmpty(product.Sizes) 
	suite.NotEmpty(product.Color) 

	suite.Contains(product.Color , int64(color.Id)) 
	suite.Contains(product.Sizes ,int64( size.Id)) 

	var products []domain.Product
	err = suite.GetProducts(&products) 
	suite.NoError(err) 
	suite.NotEmpty(products)
}

func(suite*dbsuite)TestCreateProduct_Detail(){
		var   category =  domain.Category{
		Name: "Clothing",
	}
	err:=suite.CreateCategory(&category) 

suite.NoError(err)
	var audience=  domain.Audience{
		Name: "Women", 
		
	}
 

	
  err = suite.CreateAudience(&audience) 

suite.NoError(err)

var m  ="xs|s|M|L"
  var typpe =  domain.Typpe{
 		Name: "T-shirt", 
		CategoryID: &category.Id,
		AudienceID: &audience.Id, 
		UnitOfMeasurement:&m ,
  } 

    err = suite.CreateTypes(&typpe) 

	suite.NoError(err)
	suite.NotEmpty(typpe.Id) 
	suite.Equal(typpe.AudienceID , &audience.Id) 
	suite.Equal(typpe.CategoryID , &category.Id)
 

	var brand=  domain.Brand{
		Name: "Adidas",
	}
	err=  suite.CreateBrand(&brand)
	suite.NoError(err) 
   
	var color = domain.Color{
		Name: "White",
		Code: "#fff",
	}

	err =  suite.CreateColor(&color) 
	suite.NoError(err) 

	size := domain.Sizee{
		Name: "XL",
	}

	err = suite.CreateSizee(&size) 
	suite.NoError(err) 

	product:= domain.Product{
		Name: "T-shirt",
		Color: []int64{int64(color.Id)},
		Sizes: []int64{int64(size.Id)} ,
		Quantity: 10, 
		MaxPrice: 100,
		MinPrice: 12, 
		BrandID: &brand.Id, 
		TyppeID: &typpe.Id,
	}

	err= suite.CreateProduct(&product) 
	suite.NoError(err) 

	suite.NotEmpty(product.Id) 
	suite.NotEmpty(product.Sizes) 
	suite.NotEmpty(product.Color) 

	suite.Contains(product.Color , int64(color.Id)) 
	suite.Contains(product.Sizes ,int64( size.Id)) 


	product_detail := domain.Product_Detail{
	 ProductID: &product.Id,
	ColorID: &color.Id,
	SizeeID: &size.Id,
	Price: 100,
	Quantity: 10,
	}
	err = suite.CreateProduct_Detail(&product_detail)
	suite.NoError(err) 

		product_detail1 := domain.Product_Detail{
	 ProductID: &product.Id,
	ColorID: &color.Id,
	SizeeID: &size.Id,
	Price: 100,
	Quantity: 10,
	}
	err = suite.CreateProduct_Detail(&product_detail1)
	suite.Error(err) 

	var products_details []domain.Product_Detail
	err =suite.GetProductDetails(&products_details)

	suite.NoError(err)
	suite.NotEmpty(products_details)  


}


func (suite *dbsuite)TestCreateOrder(){

 var   category =  domain.Category{
		Name: "Clothing",
	}
	err:=suite.CreateCategory(&category) 

suite.NoError(err)
	var audience=  domain.Audience{
		Name: "Women", 
		
	}
 

	
  err = suite.CreateAudience(&audience) 

suite.NoError(err)

var m  ="xs|s|M|L"
  var typpe =  domain.Typpe{
 		Name: "T-shirt", 
		CategoryID: &category.Id,
		AudienceID: &audience.Id, 
		UnitOfMeasurement:&m ,
  } 

    err = suite.CreateTypes(&typpe) 

	suite.NoError(err)
	suite.NotEmpty(typpe.Id) 
	suite.Equal(typpe.AudienceID , &audience.Id) 
	suite.Equal(typpe.CategoryID , &category.Id)
 

	var brand=  domain.Brand{
		Name: "Adidas",
	}
	err=  suite.CreateBrand(&brand)
	suite.NoError(err) 
   
	var color = domain.Color{
		Name: "White",
		Code: "#fff",
	}

	err =  suite.CreateColor(&color) 
	suite.NoError(err) 

	size := domain.Sizee{
		Name: "XL",
	}

	err = suite.CreateSizee(&size) 
	suite.NoError(err) 

	product:= domain.Product{
		Name: "T-shirt",
		Color: []int64{int64(color.Id)},
		Sizes: []int64{int64(size.Id)} ,
		Quantity: 10, 
		MaxPrice: 100,
		MinPrice: 12, 
		BrandID: &brand.Id, 
		TyppeID: &typpe.Id,
	}

	err= suite.CreateProduct(&product) 
	suite.NoError(err) 

	err= suite.CreateProduct(&product) 
	suite.NoError(err) 

	suite.NotEmpty(product.Id) 
	suite.NotEmpty(product.Sizes) 
	suite.NotEmpty(product.Color) 

	suite.Contains(product.Color , int64(color.Id)) 
	suite.Contains(product.Sizes ,int64( size.Id)) 


	product_detail := domain.Product_Detail{
	 ProductID: &product.Id,
	ColorID: &color.Id,
	SizeeID: &size.Id,
	Price: 100,
	Quantity: 10,
	}
	err = suite.CreateProduct_Detail(&product_detail)
	suite.NoError(err) 

		product_detail1 := domain.Product_Detail{
	 ProductID: &product.Id,
	ColorID: &color.Id,
	SizeeID: &size.Id,
	Price: 100,
	Quantity: 10,
	}
	err = suite.CreateProduct_Detail(&product_detail1)
	suite.Error(err) 

	var products_details []domain.Product_Detail
	err =suite.GetProductDetails(&products_details)

	suite.NoError(err)
	suite.NotEmpty(products_details)  



	var user  = &domain.User{
	Email: "fdfdfd", 
	Fullname: "ds",
	Avatar: "dsewe",
	Username: "_deew",
 }
	err= suite.CreateUser(user) 

	suite.NoError(err)
 suite.NotNil(user) 
  
     
  var  Order_details  []domain.OrderDetail 
 
  for i:=1 ; i<=2 ; i++{
 
	var id =  uint(i)
	Order_details = append(Order_details, domain.OrderDetail{
		ProductID:&id ,
		Quantity: 2,
		Price: 3,
	})
  }


    var newOrder=  domain.Orders{
		UserrID: &user.ID, 
		TotalQuantity: 10,
		TotalPrice: 20,
		OrderDetails: Order_details,
	}
   
	err  = suite.CreateOrders(&newOrder)
	 
	suite.NoError(err) 
	suite.NotEmpty(newOrder.OrderDetails)
 
	suite.T().Log(newOrder)
  var  orders []domain.Orders
	err = suite.GetOrders(&orders) 
	 
	suite.NoError(err) 
	suite.NotEmpty(newOrder.OrderDetails)
suite.Equal(newOrder.Id , uint(1)) 
suite.NotEmpty(orders[0].OrderDetails)
	suite.T().Log(orders)
 
	
	var userOrders domain.Orders =  domain.Orders{UserrID: &user.ID }
	err=  suite.GetUserOrders(&userOrders)

	suite.NoError(err)  
	suite.NotEmpty(userOrders.OrderDetails)
	

}

func(suite *dbsuite)TestWishList(){
	 var   category =  domain.Category{
		Name: "Clothing",
	}
	err:=suite.CreateCategory(&category) 

suite.NoError(err)
	var audience=  domain.Audience{
		Name: "Women", 
		
	}
 

	
  err = suite.CreateAudience(&audience) 

suite.NoError(err)

var m  ="xs|s|M|L"
  var typpe =  domain.Typpe{
 		Name: "T-shirt", 
		CategoryID: &category.Id,
		AudienceID: &audience.Id, 
		UnitOfMeasurement:&m ,
  } 

    err = suite.CreateTypes(&typpe) 

	suite.NoError(err)
	suite.NotEmpty(typpe.Id) 
	suite.Equal(typpe.AudienceID , &audience.Id) 
	suite.Equal(typpe.CategoryID , &category.Id)
 

	var brand=  domain.Brand{
		Name: "Adidas",
	}
	err=  suite.CreateBrand(&brand)
	suite.NoError(err) 
   
	var color = domain.Color{
		Name: "White",
		Code: "#fff",
	}

	err =  suite.CreateColor(&color) 
	suite.NoError(err) 

	size := domain.Sizee{
		Name: "XL",
	}

	err = suite.CreateSizee(&size) 
	suite.NoError(err) 

	product:= domain.Product{
		Name: "T-shirt",
		Color: []int64{int64(color.Id)},
		Sizes: []int64{int64(size.Id)} ,
		Quantity: 10, 
		MaxPrice: 100,
		MinPrice: 12, 
		BrandID: &brand.Id, 
		TyppeID: &typpe.Id,
	}

	err= suite.CreateProduct(&product) 
	suite.NoError(err) 

	err= suite.CreateProduct(&product) 
	suite.NoError(err) 

	suite.NotEmpty(product.Id) 
	suite.NotEmpty(product.Sizes) 
	suite.NotEmpty(product.Color) 

	suite.Contains(product.Color , int64(color.Id)) 
	suite.Contains(product.Sizes ,int64( size.Id)) 


	product_detail := domain.Product_Detail{
	 ProductID: &product.Id,
	ColorID: &color.Id,
	SizeeID: &size.Id,
	Price: 100,
	Quantity: 10,
	}
	err = suite.CreateProduct_Detail(&product_detail)
	suite.NoError(err) 

		product_detail1 := domain.Product_Detail{
	 ProductID: &product.Id,
	ColorID: &color.Id,
	SizeeID: &size.Id,
	Price: 100,
	Quantity: 10,
	}
	err = suite.CreateProduct_Detail(&product_detail1)
	suite.Error(err) 

	var products_details []domain.Product_Detail
	err =suite.GetProductDetails(&products_details)

	suite.NoError(err)
	suite.NotEmpty(products_details)  



	var user  = &domain.User{
	Email: "fdfdfd", 
	Fullname: "ds",
	Avatar: "dsewe",
	Username: "_deehw",
 }
	err= suite.CreateUser(user) 

	suite.NoError(err)
 suite.NotNil(user)
  
  var wishlist = []domain.Wishlist{
	{
		UserID: &user.ID, 
		Products: domain.Product{
			Id: product.Id,
		},
	 },
  }
     err=  suite.CreateWishlist(&wishlist)
	 suite.NoError(err) 
	  
	 userWishlist , err := suite.GetWishlistByUserID(wishlist)
  suite.NoError(err) 
 
  suite.T().Log(userWishlist)
    suite.NotEmpty(userWishlist[0].Products.Name)

}
func(suite *dbsuite)TestGetUserByEmail(){
	var user = &domain.User{Email: "denlinato@gmail.com"}
	
	err:= suite.GetUserBYEmail(user) 
	suite.T().Logf("%s" , err)
	suite.Error(err)
	suite.ErrorContains(err ,"user not found")
	suite.Empty(user.Fullname)
}
func TestRun(t *testing.T){
   suite.Run(t ,&dbsuite{})
}



