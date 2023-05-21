package handlers

import (
	dto "myveda/dto/result"
	usersdto "myveda/dto/users"
	"myveda/models"
	"myveda/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type handler struct {
	Repository repositories.UserRepository
}

func HandlerUsers(Repository repositories.UserRepository) *handler {
	return &handler{Repository}
}

func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.Repository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(200, dto.SuccessResult{Code: 200, Data: users})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.Repository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(200, dto.SuccessResult{Code: 200, Data: user})
}

// func (h *handler) CreateUser(c echo.Context) error {
// 	request := new(usersdto.CreateUser)
// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	validate := validator.New()
// 	err := validate.Struct(request)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	user := models.User{
// 		Email:    request.Email,
// 		Password: request.Password,
// 		Fullname: request.Fullname,
// 		Phone:    request.Phone,
// 		Address:  request.Address,
// 	}

// 	data, err := h.Repository.CreateUser(user)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}

// 	return c.JSON(200, dto.SuccessResult{Code: 200, Data: ConvertResponse(data)})
// }

func (h *handler) UpdateUser(c echo.Context) error {
	request := new(usersdto.UpdateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.Repository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	if request.Fullname != "" {
		user.Fullname = request.Fullname
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Address != "" {
		user.Address = request.Address
	}

	data, err := h.Repository.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(200, dto.SuccessResult{Code: 200, Data: ConvertResponse(data)})

}

func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.Repository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.Repository.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(200, dto.SuccessResult{Code: 200, Data: ConvertResponse(data)})
}

func ConvertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		Fullname: u.Fullname,
		Phone:    u.Phone,
		Address:  u.Address,
	}

}
