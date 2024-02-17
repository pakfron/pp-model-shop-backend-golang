package entities

import "github.com/golang-jwt/jwt/v5"

type UserRegisterReq struct {
	UserName string `json:"username" binding:"required" validate:"required,max=16,min=3"`
	PassWord string `json:"password" binding:"required" validate:"required,max=16,min=3"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
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
