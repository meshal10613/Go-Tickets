package user

import (
	"errors"
	httpsresponse "go-tickets/internel/httpsResponse"
	"go-tickets/internel/user/dto"
	"net/http"

	"github.com/labstack/echo/v5"
)

type userHandler struct {
	service *service
}

func NewUserHandler(service *service) *userHandler {
	return &userHandler{
		service: service,
	}
}

func (h *userHandler) RegisterUser(ctx *echo.Context) error {
	var req dto.RegisterUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpsresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpsresponse.Error{
			Success: false,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.RegisterUser(&req)
	if err != nil {
		if errors.Is(err, ErrorAlreadyExist) {
			return ctx.JSON(http.StatusConflict, httpsresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, httpsresponse.Error{
			Success: false,
			Message: "Failed to register user",
			Details: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, httpsresponse.Success{
		Success: true,
		Message: "User registered successfully",
		Data:    response,
	})
}

func (h *userHandler) LoginUser(ctx *echo.Context) error {
	var req dto.LoginUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpsresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, httpsresponse.Error{
			Success: false,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.LoginUser(&req)
	if err != nil {
		if errors.Is(err, ErrorInvalidCredentials) {
			return ctx.JSON(http.StatusUnauthorized, httpsresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, httpsresponse.Error{
			Success: false,
			Message: "Failed to login user",
			Details: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, httpsresponse.Success{
		Success: true,
		Message: "User logged in successfully",
		Data:    response,
	})
}
