package usecase

import (
	"errors"
	"lintaspay/internal/modules/transactions/domain"
	"lintaspay/internal/modules/transactions/repository"
	wallet "lintaspay/internal/modules/wallet/repository"

	"gorm.io/gorm"
)

type transactionsUseCase struct {
	db         *gorm.DB
	repo       repository.TransactionRepository
	walletRepo wallet.WalletRepository
}
type TransactionUsecase interface {
	TopUp(userID uint, amount int64, ref string) error
	Transfer(fromUserID, toUserID uint, amount int64, ref string) error
	History(walletID uint) ([]domain.Transaction, error)
}

func NewTransactionUsecase(db *gorm.DB, trxRepo repository.TransactionRepository, walletRepo wallet.WalletRepository,
) TransactionUsecase {
	return &transactionsUseCase{
		db:         db,
		repo:       trxRepo,
		walletRepo: walletRepo,
	}
}
func (u *transactionsUseCase) TopUp(userID uint, amount int64, ref string) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return u.db.Transaction(func(tx *gorm.DB) error {
		if trx, _ := u.repo.FindByReferenceID(tx, ref); trx != nil {
			return errors.New("transaction already processed")
		}

		wallet, err := u.walletRepo.FindByUserIDTx(tx, userID)
		if err != nil {
			return err
		}

		newBalance := wallet.Balance + amount
		if err := u.walletRepo.UpdateWalletTx(tx, wallet.ID, newBalance); err != nil {
			return err
		}

		return u.repo.CreateTx(tx, &domain.Transaction{
			WalletID:    wallet.ID,
			Type:        domain.Credit,
			Amount:      amount,
			ReferenceID: ref,
		})
	})
}
func (u *transactionsUseCase) Transfer(fromUserID, toUserID uint, amount int64, ref string,
) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if fromUserID == toUserID {
		return errors.New("cannot transfer to same user")
	}

	return u.db.Transaction(func(tx *gorm.DB) error {
		// idempotency
		if trx, _ := u.repo.FindByReferenceID(tx, ref); trx != nil {
			return errors.New("transaction already processed")
		}

		fromWallet, err := u.walletRepo.FindByUserIDTx(tx, fromUserID)
		if err != nil {
			return err
		}

		toWallet, err := u.walletRepo.FindByUserIDTx(tx, toUserID)
		if err != nil {
			return err
		}

		if fromWallet.Balance < amount {
			return errors.New("insufficient balance")
		}

		// update balances
		if err := u.walletRepo.UpdateWalletTx(
			tx, fromWallet.ID, fromWallet.Balance-amount,
		); err != nil {
			return err
		}

		if err := u.walletRepo.UpdateWalletTx(
			tx, toWallet.ID, toWallet.Balance+amount,
		); err != nil {
			return err
		}

		// debit transaction
		if err := u.repo.CreateTx(tx, &domain.Transaction{
			WalletID:    fromWallet.ID,
			Type:        domain.Debit,
			Amount:      amount,
			ReferenceID: ref,
		}); err != nil {
			return err
		}

		// credit transaction
		return u.repo.CreateTx(tx, &domain.Transaction{
			WalletID:    toWallet.ID,
			Type:        domain.Credit,
			Amount:      amount,
			ReferenceID: ref,
		})
	})
}
func (u *transactionsUseCase) History(walletID uint) ([]domain.Transaction, error) {
	return u.repo.FindByWalletID(walletID)
}
