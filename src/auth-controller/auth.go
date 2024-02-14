package auth

import (
	"net/http"
	"pp-model-shop-backend/validate"

	"github.com/gin-gonic/gin"
)

type register struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Register(c *gin.Context) {
	var json register

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if email := validate.ValidateRegister(json.Email); email == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email is Invalid"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"register": json})
}
