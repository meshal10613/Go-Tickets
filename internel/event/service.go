package event

import (
	"go-tickets/internel/event/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) Create(req dto.CreateRequest) (*dto.Response, error) {
	var err error
	event := Event{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		StartsAt:         req.StartsAt,
		TotalTickets:     req.TotalTickets,
		AvailableTickets: req.TotalTickets,
		Price:            req.Price,
	}

	err = s.repo.Create(&event)
	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

func (s *service) GetAll() (*[]dto.Response, error) {
	events, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []dto.Response

	for _, event := range events {
		response = append(response, *event.ToResponse())
	}

	return &response, nil
}

func (s *service) GetByID(eventId uint) (*dto.Response, error) {
	event, err := s.repo.GetByID(eventId)

	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}
