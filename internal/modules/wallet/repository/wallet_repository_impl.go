package repository

import (
	"lintaspay/internal/modules/wallet/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type walletRepository struct {
	DB *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{DB: db}
}
func (w *walletRepository) CreateWallet(data *domain.Wallet) error {
	return w.DB.Create(data).Error
}

func (w *walletRepository) FindByUserIDTx(tx *gorm.DB, userID uint) (*domain.Wallet, error) {
	var wallet domain.Wallet

	err := tx.
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("user_id = ?", userID).
		First(&wallet).
		Error

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
func (w *walletRepository) UpdateWalletTx(tx *gorm.DB, walletID uint, balance int64) error {
	return tx.
		Model(&domain.Wallet{}).
		Where("id = ?", walletID).
		Update("balance", balance).
		Error
}
