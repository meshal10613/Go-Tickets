package event

import (
	"errors"
	"go-tickets/internel/common/query"
	"go-tickets/internel/common/querybuilder"

	"gorm.io/gorm"
)

var ErrEventNotFound = errors.New("event not found")

type Repository interface {
	Create(event *Event) error
	GetAll(opts query.QueryParams) ([]*Event, int64, error)
	GetByID(eventId uint) (*Event, error)
	Update(event *Event) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(event *Event) error {
	return r.db.Create(event).Error
}

func (r *repository) GetAll(opts query.QueryParams) ([]*Event, int64, error) {
	var events []*Event
	var total int64

	qb := querybuilder.New(r.db.Model(&Event{}))

	qb.Search(opts.Search, "title", "location")
	qb.Sort(opts.SortBy, opts.Order)

	qb.DB.Count(&total)

	qb.Paginate(opts.Page, opts.Limit)

	err := qb.DB.Find(&events).Error

	return events, total, err
}

func (r *repository) GetByID(eventId uint) (*Event, error) {
	var event Event

	err := r.db.First(&event, eventId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEventNotFound
		}

		return nil, err
	}

	return &event, nil
}

func (r *repository) Update(event *Event) error {
	return r.db.Save(event).Error
}
