package authRegister

import (
	"fmt"
	"net/http"
	pp_model_schema "pp-model-shop-backend/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		// fmt.Println()

		c.JSON(http.StatusBadRequest, gin.H{"error": "input invalid "})
		return
	}

	users := pp_model_schema.User{
		UserName: json.UserName,
		PassWord: json.PassWord,
		Email:    json.Email,
	}
	// fmt.Println(users.ID)
	pp_model_schema.Instance.Create(&users)
	// if users.ID > 0 {
	// 	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Create Success", "userId": users.ID})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Create Failed"})
	// }

	c.JSON(http.StatusCreated, gin.H{"register": users})
}
