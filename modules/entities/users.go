package entities_user

import "github.com/golang-jwt/jwt/v5"

type UserRegisterReq struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required" `
	Email    string `json:"email" binding:"required"`
}

type UserRegisterValidate struct {
	UserName string `validate:"required,max=16,min=3,alphanum"`
	PassWord string `validate:"required,max=16,min=3"`
	Email    string `validate:"required,email"`
}
type UserLoginValidate struct {
	UserName string `validate:"required,max=16,min=3,alphanum"`
	PassWord string `validate:"required,max=16,min=3"`
}

type UserRegisterCase interface {
	NewUserCase(req *UserRegisterReq) (*UserRegisterRes, error)
}

type RoleType string

const (
	ADMIN = "ADMIN"
	USER  = "USER"
)

type CheckUser struct {
	UserName string
	Email    string
}

type UserRegisterDB struct {
	UserName string
	PassWord string
	Email    string
	Role     RoleType
}

type UserRegisterRes struct {
	UserName    string
	Role        RoleType
	AccessToken string
}

type Playload struct {
	UserName string
	Role     RoleType
}

type MyCustomClaims struct {
	Playload Playload
	jwt.RegisteredClaims
}

type UserLoginReq struct {
	UserName string `json:"username" binding:"required" validate:"required,max=16,min=3"`
	PassWord string `json:"password" binding:"required" validate:"required,max=16,min=3"`
}

type UserLoginRes struct {
	UserName    string
	Role        RoleType
	AccessToken string
}

type DecryptPassword struct {
	HashPassword []byte
	PassWord     []byte
}
