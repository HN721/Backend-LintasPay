package usecase

import (
	"errors"
	"lintaspay/internal/modules/users/domain"
	"lintaspay/internal/modules/users/repository"
	walletDomain "lintaspay/internal/modules/wallet/domain"
	walletRepo "lintaspay/internal/modules/wallet/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	Register(name, email, password string) error
	Login(email, password string) (*domain.User, error)
}
type authUsecase struct {
	repo   repository.UserRepository
	wallet walletRepo.WalletRepository
}

func NewAuthUsecase(repo repository.UserRepository, wallet walletRepo.WalletRepository) UserUseCase {
	return &authUsecase{repo: repo, wallet: wallet}
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

	user, err = h.repo.Register(user)
	data := walletDomain.Wallet{
		UserID:  user.ID,
		Balance: 0,
	}
	return h.wallet.CreateWallet(&data)

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
