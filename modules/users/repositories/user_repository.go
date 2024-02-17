package repositories

import (
	"errors"
	"fmt"
	"pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/server"
	databases "pp-model-shop-backend/pkg/database"
)

func Register(user *databases.User, hashPassword []byte) (*databases.User, error) {
	result := server.Instance.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)

		return nil, result.Error
	}
	return user, nil
}

func CheckUser(input *entities.UserRegisterReq) error {
	var userName databases.User

	server.Instance.Where("user_name = ? AND email =?", input.UserName, input.Email).First(&userName)
	if userName.UserName == input.UserName && userName.Email == input.Email {
		err := errors.New("UserName and Email Already Use")

		return err
	}
	server.Instance.Where("user_name = ?", input.UserName).First(&userName)
	if userName.UserName == input.UserName {
		err := errors.New("UserName Already Use")
		return err
	}
	server.Instance.Where("email = ?", input.Email).First(&userName)
	if userName.Email == input.Email {
		err := errors.New("email Already Use")
		return err
	}
	return nil
}
