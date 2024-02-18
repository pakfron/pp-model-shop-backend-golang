package UserControllers

import (
	"net/http"
	entities_user "pp-model-shop-backend/modules/entities"
	user_usecase "pp-model-shop-backend/modules/users/usecase"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Error string
}

func Register(c *gin.Context) {

	var input *entities_user.UserRegisterReq

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := user_usecase.NewUserCase(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Register": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Register": user})
}

func Login(c *gin.Context) {
	var input *entities_user.UserLoginReq

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := user_usecase.UserLoginCase(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Login": user})
}
