package repositories

import (
	"errors"
	"fmt"
	"pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/server"
	databases "pp-model-shop-backend/pkg/database"
)

func CheckProduct(input *entities.CreateProductReq) error {

	var count int64
	db := server.GetDB()
	defer server.CloseDB(db)
	db.Model(&databases.Product{}).Where("name =?", input.Name).Count(&count)

	if count != 0 {
		err := errors.New("product already exist")
		return err
	}

	return nil
}

func CreateProduct(input *databases.Product, imageUrl *entities.URLProduct) (*entities.CreateProudctRes, error) {
	db := server.GetDB()
	defer server.CloseDB(db)
	result := db.Model(databases.Product{}).Create(input)
	fmt.Println(input)
	if result.Error != nil {
		return nil, result.Error
	}

	imageProduct := databases.ImageProduct{
		ImageUrl:  imageUrl.Url,
		ProductId: input.ID,
	}

	imageProductResult := db.Model(databases.ImageProduct{}).Create(&imageProduct)
	fmt.Println(imageProduct)
	if imageProductResult.Error != nil {
		return nil, imageProductResult.Error
	}

	productRes := entities.CreateProudctRes{
		Name:     input.Name,
		Series:   input.Series,
		Detail:   input.Detail,
		Type:     entities.Type(input.Type),
		Price:    input.Price,
		ImageUrl: imageProduct.ImageUrl,
	}

	return &productRes, nil
}
