package routes

import (
	productController "pp-model-shop-backend/modules/products/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoute(router *gin.RouterGroup) {

	productRouter := router.Group("/product")
	{
		productRouter.POST("/create", productController.AddProduct)
	}

}
