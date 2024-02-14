package authRegister

import (
	"fmt"
	"net/http"
	pp_model_schema "pp-model-shop-backend/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type InputRegister struct {
	UserName string `json:"username" binding:"required" validate:"required,max=16,min=3"`
	PassWord string `json:"password" binding:"required" validate:"required,max=16,min=3"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
}

type User struct {
	user_name string
	pass_word string
	email     string
}

func Register(c *gin.Context) {
	var json InputRegister

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(json); err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "input invalid "})
		return
	}

	// users := pp_model_schema.User{
	// 	UserName: json.UserName,
	// 	PassWord: json.PassWord,
	// 	Email:    json.Email,
	// }

	userExist := pp_model_schema.User{}

	pp_model_schema.Instance.Where("user_name = ?", json.UserName).First(&userExist)
	if userExist.UserName == json.UserName {
		fmt.Println(userExist.UserName)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already Exist"})
		return
	}

	pp_model_schema.Instance.Where("email =?", json.Email).First(&userExist)
	if userExist.Email == json.Email {
		fmt.Println(userExist.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(json.PassWord), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"register": hashPassword})
}
