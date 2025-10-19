// domain/user/user.go
package user

import (
	"context"
	"errors"
)

// Service (opsional)
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) RegisterUser(ctx context.Context, user *User) error {
	// Contoh business logic
	existing, _ := s.repo.FindByEmail(ctx, user.Email)
	if existing != nil {
		return errors.New("user with this email already exists")
	}
	return s.repo.Create(ctx, user)
}
