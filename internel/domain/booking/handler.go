package booking

import (
	"errors"
	"go-tickets/internel/domain/booking/dto"
	"go-tickets/internel/domain/event"
	httpresponse "go-tickets/internel/httpResponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(s *service) *handler {
	return &handler{service: s}
}

func getCurrentUserID(c *echo.Context) (uint, bool) {
	userId, ok := c.Get("user_id").(uint)
	return userId, ok
}

func bookingErrorResponse(c *echo.Context, err error) error {
	if errors.Is(err, ErrBookingNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Success: false,
			Message: "Booking not found",
		})
	}

	if errors.Is(err, event.ErrEventNotFound) {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Success: false,
			Message: "Event not found",
		})
	}

	if errors.Is(err, ErrNotEnoughTickets) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Success: false,
			Message: "Not enough tickets available",
		})
	}

	if errors.Is(err, ErrBookingAlreadyCancelled) {
		return c.JSON(http.StatusConflict, httpresponse.Error{
			Success: false,
			Message: "Booking is already cancelled",
		})
	}

	if errors.Is(err, ErrForbiddenBookingAccess) {
		return c.JSON(http.StatusForbidden, httpresponse.Error{
			Success: false,
			Message: "You do not own this booking",
		})
	}

	return c.JSON(http.StatusInternalServerError, httpresponse.Error{
		Success: false,
		Message: "Something went wrong",
		Details: err.Error(),
	})
}

func (h *handler) Create(c *echo.Context) error {
	userId, ok := getCurrentUserID(c)
	if !ok {
		return c.JSON(http.StatusUnauthorized, httpresponse.Success{
			Success: false,
			Message: "Unauthorized Access",
		})
	}

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

	response, err := h.service.Create(userId, req)

	if err != nil {
		return bookingErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Message: "Booking created successfully",
		Data: response,
	})
}