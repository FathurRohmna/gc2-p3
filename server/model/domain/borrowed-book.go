package domain

import (
	"time"

	"gorm.io/gorm"
)

type BorrowedBook struct {
	ID           string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	BookID       string    `gorm:"index"`
	UserID       string    `gorm:"index"`
	BorrowedDate time.Time `gorm:"not null"`
	ReturnDate   time.Time `gorm:"default:null"`

	CreatedAt time.Time       `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time       `gorm:"type:timestamp;default:current_timestamp;not null" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
