package querybuilder

import (
	"strings"

	"gorm.io/gorm"
)

type Builder struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Builder {
	return &Builder{DB: db}
}

func (b *Builder) Search(keyword string, columns ...string) *Builder {
	if keyword == "" {
		return b
	}

	var conditions []string
	var values []interface{}

	for _, col := range columns {
		conditions = append(conditions, col+" ILIKE ?")
		values = append(values, "%"+keyword+"%")
	}

	b.DB = b.DB.Where(strings.Join(conditions, " OR "), values...)

	return b
}

func (b *Builder) Sort(sortBy, order string) *Builder {
	if sortBy == "" {
		sortBy = "created_at"
	}

	if order != "asc" {
		order = "desc"
	}

	b.DB = b.DB.Order(sortBy + " " + order)

	return b
}

func (b *Builder) Paginate(page, limit int) *Builder {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	b.DB = b.DB.Offset(offset).Limit(limit)

	return b
}
