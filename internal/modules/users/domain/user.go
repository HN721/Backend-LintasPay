package domain

import (
	"lintaspay/internal/modules/wallet/domain"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`

	Wallet    domain.Wallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
