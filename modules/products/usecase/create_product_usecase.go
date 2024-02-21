package product_usecase

import (
	"errors"
	"pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/products/repositories"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context, input *entities.CreateProductReq) (*entities.CreateProudctRes, error) {

	if err := repositories.CheckProduct(input); err != nil {
		return nil, err
	}
	err := errors.New("test Error")

	return nil, err
}
