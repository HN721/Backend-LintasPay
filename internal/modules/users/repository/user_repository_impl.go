package repository

import (
	"errors"
	"lintaspay/internal/modules/users/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}
func (u *userRepository) Register(data *domain.User) error {
	return u.DB.Create(data).Error
}
func (u *userRepository) FindEmail(email string) (*domain.User, error) {
	var user domain.User

	err := u.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
