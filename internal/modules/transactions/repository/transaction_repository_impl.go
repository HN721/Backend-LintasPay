package repository

import (
	"lintaspay/internal/modules/transactions/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTx(tx *gorm.DB, trx *domain.Transaction) error {
	return tx.Create(trx).Error
}

func (r *transactionRepository) FindByReferenceID(tx *gorm.DB, ref string) (*domain.Transaction, error) {
	var trx domain.Transaction
	err := tx.Where("reference_id = ?", ref).First(&trx).Error
	if err != nil {
		return nil, err
	}
	return &trx, nil
}

func (r *transactionRepository) FindByWalletID(walletID uint) ([]domain.Transaction, error) {
	var trxs []domain.Transaction
	err := r.db.
		Where("wallet_id = ?", walletID).
		Order("created_at desc").
		Find(&trxs).Error

	return trxs, err
}
