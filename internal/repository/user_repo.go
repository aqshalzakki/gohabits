package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	domainUser "gohabits/internal/domain/user"
	sharedErrors "gohabits/internal/shared"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domainUser.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *domainUser.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domainUser.User, error) {
	var user domainUser.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sharedErrors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*domainUser.User, error) {
	var user domainUser.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, sharedErrors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
