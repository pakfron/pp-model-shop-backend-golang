package productController

import (
	"net/http"
	"pp-model-shop-backend/modules/entities"
	product_usecase "pp-model-shop-backend/modules/products/usecase"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var input *entities.CreateProductReq

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	}

	product, err := product_usecase.CreateProduct(c, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"addProduct": product})

}
