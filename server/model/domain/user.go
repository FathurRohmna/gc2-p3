package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`

	CreatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
