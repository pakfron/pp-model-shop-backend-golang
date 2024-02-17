package UserControllers

import (
	"net/http"
	"pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/users/usecase"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Error string
}

func Register(c *gin.Context) {

	var input *entities.UserRegisterReq

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := usecase.NewUserCase(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Register": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Register": user})
}
