package event

import "go-tickets/internel/event/dto"

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
