package container

import (
	walletHandler "lintaspay/internal/modules/wallet/handler"
	walletRepo "lintaspay/internal/modules/wallet/repository"
	walletUseCase "lintaspay/internal/modules/wallet/usecase"

	"gorm.io/gorm"
)

type WalletContainer struct {
	WalletHandler *walletHandler.WalletHandler
}

func NewWalletContainer(db *gorm.DB) *WalletContainer {
	repo := walletRepo.NewWalletRepository(db)
	uc := walletUseCase.NewWalletUsecase(repo)
	handler := walletHandler.NewWalletHandler(uc)

	return &WalletContainer{
		WalletHandler: handler,
	}
}
