package handlers

import (
	"context"
	"fmt"
	productdto "myveda/dto/product"
	dto "myveda/dto/result"
	"myveda/models"
	"myveda/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerProduct struct {
	product repositories.ProductRepository
}

func ProductRepo(product repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{product}
}

func (h *handlerProduct) FindProduct(c echo.Context) error {
	products, err := h.product.FindProduct()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(200, dto.SuccessResult{Code: 200, Data: products})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)

	price, _ := strconv.Atoi(c.FormValue("price"))

	request := productdto.CreateProduct{
		Title:     c.FormValue("title"),
		Thumbnail: c.FormValue("thumbnail"),
		Image:     dataFile,
		Price:     price,
		Size:      c.FormValue("size"),
		Style:     c.FormValue("style"),
		Material:  c.FormValue("material"),
		Color:     c.FormValue("color"),
		Condition: c.FormValue("condition"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "validation error!"})
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "myveda"})
	if err != nil {
		fmt.Println(err.Error())
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	Products := models.Product{
		Title:     request.Title,
		Thumbnail: request.Thumbnail,
		Image:     resp.SecureURL,
		Price:     request.Price,
		Size:      request.Size,
		Style:     request.Style,
		Material:  request.Material,
		Color:     request.Color,
		Condition: request.Condition,
		UserID:    int(userId),
	}

	data, err := h.product.CreateProduct(Products)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "data not found!"})
	}

	data, _ = h.product.GetProduct(data.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
