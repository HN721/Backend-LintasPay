package domain

import "time"

type Wallet struct {
	ID      uint  `gorm:"primaryKey"`
	UserID  uint  `gorm:"uniqueIndex;not null"`
	Balance int64 `gorm:"not null;default:0"`
	// Transactions []Transaction `gorm:"foreignKey:WalletID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
