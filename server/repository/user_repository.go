package repository

import (
	"context"
	"server/model/domain"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	FindByUsername(ctx context.Context, tx *gorm.DB, username string) (domain.User, error)
}
