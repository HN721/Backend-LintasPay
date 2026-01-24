package repository

import "lintaspay/internal/modules/wallet/domain"

type WalletRepository interface {
	CreateWallet(data *domain.Wallet) error
	UpdateWallet(data *domain.Wallet) error
}
