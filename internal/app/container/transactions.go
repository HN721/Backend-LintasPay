package container

import (
	trxHandler "lintaspay/internal/modules/transactions/handler"
	trxRepo "lintaspay/internal/modules/transactions/repository"
	walletRepo "lintaspay/internal/modules/wallet/repository"

	trxUseCase "lintaspay/internal/modules/transactions/usecase"

	"gorm.io/gorm"
)

type TrxContainer struct {
	TrxHandler *trxHandler.TransactionHandler
}

func NewTrxContainer(db *gorm.DB) *TrxContainer {
	repo := trxRepo.NewTransactionRepository(db)
	walletRepo := walletRepo.NewWalletRepository(db)

	uc := trxUseCase.NewTransactionUsecase(db, repo, walletRepo)
	handler := trxHandler.NewTransactionHandler(uc)

	return &TrxContainer{
		TrxHandler: handler,
	}
}
