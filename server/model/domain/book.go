package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Title         string    `gorm:"not null" json:"title"`
	Author        string    `gorm:"not null" json:"author"`
	PublishedDate time.Time `gorm:"not null" json:"published_date"`
	Status        string    `gorm:"not null;default:Available" json:"status"`

	CreatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp;not null" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
