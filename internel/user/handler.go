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

func (h *userHandler) CreateUser(ctx *echo.Context) error {
	var req dto.CreateUserRequest
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

	response, err := h.service.CreateUser(&req)
	if err != nil {
		if errors.Is(err, ErrorAlreadyExist) {
			return ctx.JSON(http.StatusConflict, httpsresponse.Error{
				Success: false,
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, httpsresponse.Error{
			Success: false,
			Message: "Failed to create user",
			Details: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, httpsresponse.Success{
		Success: true,
		Message: "User created successfully",
		Data:    response,
	})
}
