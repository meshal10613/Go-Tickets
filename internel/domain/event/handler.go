package event

import (
	"errors"
	"go-tickets/internel/common/query"
	"go-tickets/internel/domain/event/dto"
	httpresponse "go-tickets/internel/httpResponse"
	"net/http"
	"strconv"

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

	response, err := h.service.Create(req)
	if err != nil {
		return eventErrorResponse(c, err)
	}

	return c.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Message: "Event created successfully",
		Data:    response,
	})
}

func (h *handler) GetAll(c *echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	events, err := h.service.GetAll(query.QueryParams{
		Page:   page,
		Limit:  limit,
		Search: c.QueryParam("search"),
		SortBy: c.QueryParam("sortBy"),
		Order:  c.QueryParam("order"),
	})
	if err != nil {
		return eventErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, httpresponse.Success{
		Success: true,
		Message: "Events retrived successfully",
		Data:    events.Data,
		Meta:    events.Meta,
	})
}

func (h *handler) GetByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}

	response, err := h.service.GetByID(uint(id))
	if err != nil {
		return eventErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, httpresponse.Success{
		Success: true,
		Message: "Event retrived successfully",
		Data:    response,
	})
}

func (h *handler) Update(c *echo.Context) error {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Success: false,
			Message: "Invalid event id",
			Details: err.Error(),
		})
	}

	var req dto.UpdateRequest

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

	response, err := h.service.Update(uint(eventId), req)
	if err != nil {
		return eventErrorResponse(c, err)
	}

	return c.JSON(http.StatusOK, httpresponse.Success{
		Success: false,
		Message: "Event updated successfully",
		Data:    response,
	})
}
