package booking

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"go-tickets/internel/domain/booking/dto"
	"go-tickets/internel/domain/event"
	"time"
)

type service struct {
	bookingRepo Repository
	eventRepo   event.Repository
}

func NewService(bookingRepo Repository, eventRepo event.Repository) *service {
	return &service{
		bookingRepo: bookingRepo,
		eventRepo:   eventRepo,
	}
}

func generateBookingCode() string {
	randomBytes := make([]byte, 4)

	if _, err := rand.Read(randomBytes); err != nil {
		return fmt.Sprintf("EVT-%d", time.Now().UnixNano())
	}

	return fmt.Sprintf(
		"BK-%s-%s",
		time.Now().Format("20060102150405"),
		hex.EncodeToString(randomBytes),
	)
}

func (s *service) Create(userId uint, req dto.CreateRequest) (*dto.Response, error) {
	event, err := s.eventRepo.GetByID(req.EventID)
	if err != nil {
		return nil, err
	}

	if event.AvailableTickets < req.Quantity {
		return nil, ErrNotEnoughTickets
	}

	booking := &Booking{
		UserID:      userId,
		EventID:     req.EventID,
		Quantity:    req.Quantity,
		TotalPrice:  req.Quantity * event.Price,
		Status:      BookingConfirmed,
		BookingCode: generateBookingCode(),
	}
	err = s.bookingRepo.Create(booking)
	if err != nil {
		return nil, err
	}

	event.AvailableTickets = event.AvailableTickets - req.Quantity
	if err := s.eventRepo.Update(event); err != nil {
		return nil, err
	}

	return booking.ToResponse(), nil
}

func (s *service) GetMyBookings(userId uint, req dto.CreateRequest) ([]*dto.Response, error) {
	bookings, err := s.bookingRepo.GetByUserID(userId)
	if err != nil {
		return nil, err
	}

	response := make([]*dto.Response, len(bookings)) //? Initialize the slice with the correct length
	for i, b := range bookings {
		response[i] = b.ToResponse()
	}

	return response, nil
}
