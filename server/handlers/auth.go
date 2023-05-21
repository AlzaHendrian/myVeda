package handlers

import (
	"log"
	authdto "myveda/dto/auth"
	dto "myveda/dto/result"
	"myveda/models"
	"myveda/pkg/bycript"
	jwtToken "myveda/pkg/jwt"
	"myveda/repositories"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepo repositories.RepoAuth
}

func HandlerAuth(AuthRepo repositories.RepoAuth) *handlerAuth {
	return &handlerAuth{AuthRepo}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.Register)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Register Bind Failed!"})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Regist Validation Failed!"})
	}

	password, err := bycript.HashingPass(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Pass Regist not Found!"})
	}

	user := models.User{
		Email:    request.Email,
		Password: password,
		Fullname: request.Fullname,
		Phone:    request.Phone,
		Address:  request.Address,
		Role:     "customer",
	}

	data, err := h.AuthRepo.Register(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "data of user not found!"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = data.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	registerResponse := authdto.RespRegister{
		Email: data.Email,
		Token: token,
	}
	return c.JSON(200, dto.SuccessResult{Code: 200, Data: registerResponse})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.Login)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "request login not found!"})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepo.Login(user.Email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "data of user not found!"})
	}

	isValid := bycript.CheckPassHash(request.Password, user.Password)

	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Wrong Email or Password!"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.RespLogin{
		Email: user.Email,
		Token: token,
		Role:  user.Role,
		ID:    user.ID,
	}

	return c.JSON(200, dto.SuccessResult{Code: 200, Data: loginResponse})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepo.CheckAuth(int(userId))

	return c.JSON(200, dto.SuccessResult{Code: 200, Data: user})
}
