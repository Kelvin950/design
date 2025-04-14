package ports

import "github.com/kelvin950/desing/internals/application/domain"

type Firebase interface{
	Createuser(user *domain.User)error
	LoginUser(user *domain.User)(string , error)
}