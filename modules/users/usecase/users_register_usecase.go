package user_usecase

import (
	"os"
	entities "pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/users/repositories"
	databases "pp-model-shop-backend/pkg/database"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func NewUserCase(req *entities.UserRegisterReq) (*entities.UserRegisterRes, error) {

	err := repositories.CheckCreateUser(req)
	if err != nil {
		return nil, err
	}

	if err := validateInputRes(req); err != nil {
		return nil, err
	}

	hashPassword, err := HashPassword(req.PassWord)
	if err != nil {
		return nil, err
	}
	inputDB := inputRegDB(req, hashPassword)

	user, err := repositories.Register(&inputDB, hashPassword)
	if err != nil {
		return nil, err
	}
	accessToken, err := createToken(user)
	if err != nil {
		return nil, err
	}

	output := RegisterRespone(user, accessToken)
	return output, nil
}

func validateInputRes(input *entities.UserRegisterReq) error {

	userRegisterRequest := entities.UserRegisterValidate{
		UserName: input.UserName,
		PassWord: input.PassWord,
		Email:    input.Email,
	}

	validate := validator.New()
	if err := validate.Struct(&userRegisterRequest); err != nil {
		return err
	}
	return nil
}

func HashPassword(input string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 10)
	return bytes, err
}

func inputRegDB(input *entities.UserRegisterReq, hashPassword []byte) databases.User {

	inputRegForDB := databases.User{
		UserName: input.UserName,
		PassWord: string(hashPassword),
		Email:    input.Email,
	}
	return inputRegForDB

}

func createToken(input *databases.User) (*string, error) {

	Playload := entities.Playload{
		UserName: input.UserName,
		Role:     entities.RoleType(input.Role),
	}

	claim := entities.MyCustomClaims{
		Playload: Playload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	godotenv.Load("../../../.env")
	secret_key := os.Getenv("SECRET_KEY")
	var mySigningKey = []byte(secret_key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	accessToken, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}
	return &accessToken, nil
}

func RegisterRespone(input *databases.User, AccessToken *string) *entities.UserRegisterRes {

	output := entities.UserRegisterRes{
		UserName:    input.UserName,
		Role:        entities.RoleType(input.Role),
		AccessToken: *AccessToken,
	}

	return &output

}
