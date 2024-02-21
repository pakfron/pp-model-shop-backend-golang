package entities

import "mime/multipart"

type Type string

const (
	Nendoroid = "Nendoroid"
	Figma     = "Figma"
	Figure    = "Figure"
)

type CreateProductReq struct {
	Name   string                `form:"name" binding:"required"`
	Series string                `form:"series" binding:"required"`
	Detail string                `form:"detail" binding:"required"`
	Price  float32               `form:"price" binding:"required"`
	Type   Type                  `form:"type" binding:"required"`
	Image  *multipart.FileHeader `form:"image" binding:"required"`
}

type CreateProudctRes struct {
	Name   string
	Series string
	Detail string
	Price  float32
	Type   Type
}
