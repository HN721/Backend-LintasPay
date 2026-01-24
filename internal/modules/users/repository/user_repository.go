package repository

import "lintaspay/internal/modules/users/domain"

type UserRepository interface {
	Register(data *domain.User) error
	FindEmail(email string) (*domain.User, error)
}
