package repositories

import (
	"errors"
	"fmt"
	entities "pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/server"
	databases "pp-model-shop-backend/pkg/database"
)

func Register(user *databases.User, hashPassword []byte) (*databases.User, error) {

	db := server.GetDB()
	defer server.CloseDB(db)
	result := db.Create(&user)
	if result.Error != nil {

		return nil, result.Error
	}

	return user, nil
}

func CheckCreateUser(input *entities.UserRegisterReq) error {

	var count int64
	db := server.GetDB()
	defer server.CloseDB(db)
	db.Model(&databases.User{}).Where("user_name = ? AND email =?", input.UserName, input.Email).Count(&count)
	if count != 0 {
		fmt.Println(count)
		err := errors.New("UserName and Email Already Use")
		return err
	}

	db.Model(&databases.User{}).Where("user_name = ?", input.UserName).Count(&count)
	if count > 0 {
		err := errors.New("UserName Already Use")
		return err
	}
	db.Model(&databases.User{}).Where("email = ?", input.Email).Count(&count)
	if count > 0 {
		err := errors.New("email Already Use")
		return err
	}
	return nil
}

func CheckUserLogin(input *entities.UserLoginReq) (*databases.User, error) {

	var user *databases.User
	db := server.GetDB()
	defer server.CloseDB(db)
	result := db.Model(&databases.User{}).Where("user_name =?", input.UserName).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil

}
