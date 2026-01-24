package repository

import (
	"lintaspay/internal/modules/wallet/domain"

	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(data *domain.Wallet) error
	FindByUserIDTx(tx *gorm.DB, user_id uint) (*domain.Wallet, error)
	UpdateWalletTx(tx *gorm.DB, walletID uint, balance int64) error
}
