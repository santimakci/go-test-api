package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User holds personal user information
type Base struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"type:datetime;not null;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"type:datetime;not null;default:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Base) GoString() string {
	return `
{
	ID: %s,
	CreatedAt: %s,
	UpdatedAt: %s,
	DeletedAt: %s,
}
`
}

// This functions are called before creating Base
func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}
