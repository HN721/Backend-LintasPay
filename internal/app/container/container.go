package container

import "gorm.io/gorm"

type Container struct {
	Wallet *WalletContainer
	User   *UserContainer
	Trx    *TrxContainer
}

func NewContainer(db *gorm.DB) *Container {
	return &Container{
		Wallet: NewWalletContainer(db),
		User:   NewUserContainer(db),
		Trx:    NewTrxContainer(db),
	}
}
