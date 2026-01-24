package domain

import (
	"lintaspay/internal/modules/transactions/domain"
	"time"
)

type Wallet struct {
	ID           uint                 `gorm:"primaryKey"`
	UserID       uint                 `gorm:"uniqueIndex;not null"`
	Balance      int64                `gorm:"not null;default:0"`
	Transactions []domain.Transaction `gorm:"foreignKey:WalletID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
