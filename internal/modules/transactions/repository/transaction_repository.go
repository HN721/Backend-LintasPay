package repository

import (
	"lintaspay/internal/modules/transactions/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTx(tx *gorm.DB, trx *domain.Transaction) error
	FindByReferenceID(tx *gorm.DB, referenceID string) (*domain.Transaction, error)
	FindByWalletID(walletID uint) ([]domain.Transaction, error)
}
