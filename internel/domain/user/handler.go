package user

import (
	"errors"
	httpresponse "go-tickets/internel/httpResponse"
	"go-tickets/internel/domain/user/dto"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) RegisterUser(ctx *echo.Context) error {
	var req dto.RegisterUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.RegisterUser(&req)
	if err != nil {
		if errors.Is(err, ErrorAlreadyExist) {
			return ctx.JSON(http.StatusConflict, httpresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, httpresponse.Error{
			Success: false,
			Message: "Failed to register user",
			Details: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Message: "User registered successfully",
		Data:    response,
	})
}

func (h *handler) LoginUser(ctx *echo.Context) error {
	var req dto.LoginUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(&req)
	if err != nil {
		if errors.Is(err, ErrorInvalidCredentials) {
			return ctx.JSON(http.StatusUnauthorized, httpresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, httpresponse.Error{
			Success: false,
			Message: "Failed to login user",
			Details: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Message: "User logged in successfully",
		Data:    response,
	})
}

func (h *handler) GetMe(ctx *echo.Context) error {
	userID, ok := ctx.Get("user_id").(uint)
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, httpresponse.Error{
			Success: false,
			Message: "Failed to retrieve authenticated user information",
		})
	}

	email, _ := ctx.Get("email").(string)
	name, _ := ctx.Get("name").(string)

	response := dto.UserResponse{
		ID:    userID,
		Name:  name,
		Email: email,
	}

	return ctx.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Message: "User profile retrieved successfully",
		Data:    response,
	})
}
