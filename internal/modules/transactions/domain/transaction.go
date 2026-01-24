package domain

import "time"

type TransactionType string

const (
	Credit TransactionType = "CREDIT"
	Debit  TransactionType = "DEBIT"
)

type Transaction struct {
	ID          uint            `gorm:"primaryKey"`
	WalletID    uint            `gorm:"index;not null"`
	Type        TransactionType `gorm:"type:varchar(10);not null"`
	Amount      int64           `gorm:"not null"`
	ReferenceID string          `gorm:"type:varchar(100);index"`
	CreatedAt   time.Time       `gorm:"autoCreateTime"`
}
