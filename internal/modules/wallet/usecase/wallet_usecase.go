package usecase

import (
	"errors"
	"lintaspay/internal/modules/wallet/domain"
	"lintaspay/internal/modules/wallet/repository"
)

type WalletUsecase interface {
	Create(data *domain.Wallet) error
	Update(data *domain.Wallet) error
}
type walletUsecase struct {
	repo repository.WalletRepository
}

func NewWalletUsecase(repo repository.WalletRepository) WalletUsecase {
	return &walletUsecase{
		repo: repo,
	}
}
func (u *walletUsecase) Create(data *domain.Wallet) error {
	if data.UserID == 0 {
		return errors.New("user id is required")
	}

	if data.Balance < 0 {
		return errors.New("initial balance cannot be negative")
	}

	return u.repo.CreateWallet(data)
}
func (u *walletUsecase) Update(data *domain.Wallet) error {
	if data.ID == 0 {
		return errors.New("wallet id is required")
	}

	if data.Balance < 0 {
		return errors.New("balance cannot be negative")
	}

	return u.repo.UpdateWallet(data)
}
