package container

import (
	userDelivery "lintaspay/internal/modules/users/handler"
	userRepo "lintaspay/internal/modules/users/repository"
	userUsecase "lintaspay/internal/modules/users/usecase"

	"gorm.io/gorm"
)

type UserContainer struct {
	AuthHandler *userDelivery.UserHandler
}

func NewUserContainer(db *gorm.DB) *UserContainer {
	repo := userRepo.NewUserRepository(db)
	uc := userUsecase.NewAuthUsecase(repo)
	handler := userDelivery.NewUserHandler(uc)

	return &UserContainer{
		AuthHandler: handler,
	}
}
