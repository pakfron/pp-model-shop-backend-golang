package repositories

import (
	"errors"
	"pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/server"
	databases "pp-model-shop-backend/pkg/database"
)

func CheckProduct(input *entities.CreateProductReq) error {

	var count int64

	server.Instance.Model(&databases.Product{}).Where("name =?", input.Name).Count(&count)

	if count != 0 {
		err := errors.New("product already exist")
		return err
	}

	return nil
}
