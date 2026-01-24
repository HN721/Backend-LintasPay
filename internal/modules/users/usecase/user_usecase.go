package usecase

import (
	"errors"
	"lintaspay/internal/modules/users/domain"
	"lintaspay/internal/modules/users/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(name, email, password string) error
	Login(email, password string) (*domain.User, error)
}
type authUsecase struct {
	repo repository.UserRepository
}

func NewAuthUsecase(repo repository.UserRepository) UserUseCase {
	return &authUsecase{repo: repo}
}
func (h *authUsecase) Register(name, email, password string) error {
	existing, err := h.repo.FindEmail(email)
	if err != nil && existing != nil {
		return errors.New("email already registered")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	err = h.repo.Register(user)
	return err
}
func (h *authUsecase) Login(email, password string) (*domain.User, error) {
	user, err := h.repo.FindEmail(email)

	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, err
	}
	return user, nil

}
