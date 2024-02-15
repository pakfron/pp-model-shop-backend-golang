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

type Playload struct {
	UserName string
	Role     pp_model_schema.RoleType
}

func Register(c *gin.Context) {
	var input InputRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		fmt.Println(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": "input invalid "})
		return
	}

	//check Username
	username := checkUserName(input.UserName)
	if username == input.UserName {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already Exist"})
		return
	}

	//check Email
	email := checkEmail(input.Email)
	if email == input.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already Exist"})
		return
	}
	//Hash Password
	password, err := HashPassword(input.PassWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash Error"})
		return
	}

	user := pp_model_schema.User{
		UserName: input.UserName,
		PassWord: password,
		Email:    input.Email,
	}

	pp_model_schema.Instance.Create(&user)

	users := Playload{
		UserName: user.UserName,
		Role:     pp_model_schema.RoleType(user.Role),
	}
	// bytes, err := json.Marshal(users)

	// playload := string(bytes)

	c.JSON(http.StatusCreated, gin.H{"register": users})
	return
}

func checkUserName(data string) string {
	userNameDup := pp_model_schema.User{}

	pp_model_schema.Instance.Where("user_name = ?", data).First(&userNameDup)
	if userNameDup.UserName == data {
		return userNameDup.UserName
	}

	return "not Duplicate"
}

func checkEmail(data string) string {

	emailDup := pp_model_schema.User{}

	pp_model_schema.Instance.Where("email = ?", data).First(&emailDup)
	if emailDup.Email == data {
		return emailDup.Email
	}
	return "not Duplicate"

}

func HashPassword(data string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), 10)
	return string(bytes), err
}
