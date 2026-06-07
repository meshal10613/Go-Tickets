package event

import (
	"errors"
	"go-tickets/internel/event/dto"
	httpresponse "go-tickets/internel/httpResponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{
		service: s,
	}
}

func eventErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrEventNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Success: false,
			Message: "Event not found",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Success: true,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) Create(c *echo.Context) error {
	var req dto.CreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.Create(req); if err != nil {
		return eventErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Message: "Event created successfully",
		Data: response,
	})
}
