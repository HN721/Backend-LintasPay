package repository

import (
	"lintaspay/internal/modules/wallet/domain"

	"gorm.io/gorm"
)

type walletRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{DB: db}
}
func (w *walletRepository) CreateWallet(data *domain.Wallet) error {
	return w.DB.Create(&data).Error
}
func (w *walletRepository) UpdateWallet(data *domain.Wallet) error {
	return w.DB.
		Model(&domain.Wallet{}).
		Where("id = ?", data.ID).
		Updates(map[string]interface{}{
			"balance": data.Balance,
		}).Error
}
