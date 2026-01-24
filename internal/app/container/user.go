package container

import (
	userDelivery "lintaspay/internal/modules/users/handler"
	userRepo "lintaspay/internal/modules/users/repository"
	userUsecase "lintaspay/internal/modules/users/usecase"
	walletRepo "lintaspay/internal/modules/wallet/repository"

	"gorm.io/gorm"
)

type UserContainer struct {
	AuthHandler *userDelivery.UserHandler
}

func NewUserContainer(db *gorm.DB) *UserContainer {
	repo := userRepo.NewUserRepository(db)
	wallet := walletRepo.NewWalletRepository(db)
	uc := userUsecase.NewAuthUsecase(repo, wallet)
	handler := userDelivery.NewUserHandler(uc)

	return &UserContainer{
		AuthHandler: handler,
	}
}
