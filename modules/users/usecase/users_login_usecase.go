package user_usecase

import (
	"errors"
	entities_user "pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/users/repositories"
	databases "pp-model-shop-backend/pkg/database"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func UserLoginCase(input *entities_user.UserLoginReq) (*entities_user.UserLoginRes, error) {

	user, err := repositories.CheckUserLogin(input)
	if err != nil {
		return nil, err
	}

	if err := validateInputLogin(input); err != nil {
		return nil, err
	}

	password := entities_user.DecryptPassword{
		HashPassword: []byte(user.PassWord),
		PassWord:     []byte(input.PassWord),
	}

	compare := DecryptPassword(&password)

	if compare != nil {
		compare := errors.New("wrong Password")
		return nil, compare
	}

	accessToken, err := createToken(user)
	if err != nil {
		return nil, err
	}

	output := LoginRespone(user, accessToken)
	return output, nil

}

func validateInputLogin(input *entities_user.UserLoginReq) error {

	userRegisterRequest := entities_user.UserLoginValidate{
		UserName: input.UserName,
		PassWord: input.PassWord,
	}

	validate := validator.New()
	if err := validate.Struct(&userRegisterRequest); err != nil {
		return err
	}
	return nil
}

func DecryptPassword(input *entities_user.DecryptPassword) error {

	err := bcrypt.CompareHashAndPassword(input.HashPassword, input.PassWord)

	if err != nil {
		return err
	}
	return nil

}

func LoginRespone(input *databases.User, AccessToken *string) *entities_user.UserLoginRes {

	output := entities_user.UserLoginRes{
		UserName:    input.UserName,
		Role:        entities_user.RoleType(input.Role),
		AccessToken: *AccessToken,
	}
	return &output

}
