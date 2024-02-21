package product_usecase

import (
	"fmt"
	"os"
	"pp-model-shop-backend/modules/entities"
	"pp-model-shop-backend/modules/products/repositories"
	databases "pp-model-shop-backend/pkg/database"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CreateProduct(c *gin.Context, input *entities.CreateProductReq) (*entities.CreateProudctRes, error) {

	if err := repositories.CheckProduct(input); err != nil {
		return nil, err
	}

	if err := c.SaveUploadedFile(input.Image, ("../assets/image" + input.Image.Filename)); err != nil {
		return nil, err
	}
	Image, errUpload := UploadImage(c)
	if errUpload != nil {
		return nil, errUpload
	}

	product := databases.Product{
		Name:   input.Name,
		Series: input.Series,
		Detail: input.Detail,
		Price:  input.Price,
		Type:   databases.Type(input.Type),
	}

	productRes, err := repositories.CreateProduct(&product, Image)
	if err != nil {
		return nil, err
	}
	fmt.Println(Image)

	return productRes, nil
}

func UploadImage(c *gin.Context) (*entities.URLProduct, error) {

	godotenv.Load("../.env")
	CLOUND_NAME := os.Getenv("CLOUND_NAME")
	API_CLOUND := os.Getenv("API_CLOUND")
	API_CLOUND_SECRET := os.Getenv("API_CLOUND_SECRET")

	fileHeader, _ := c.FormFile("image")
	file, _ := fileHeader.Open()

	cld, _ := cloudinary.NewFromParams(CLOUND_NAME, API_CLOUND, API_CLOUND_SECRET)
	result, err := cld.Upload.Upload(c, file, uploader.UploadParams{})
	if err != nil {
		return nil, err
	}
	url := result.SecureURL
	Image := entities.URLProduct{
		Url: url,
	}

	return &Image, nil
}
