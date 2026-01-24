package main

import (
	"lintaspay/config"
	"lintaspay/internal/app/container"
	router "lintaspay/internal/app/route"
	trx "lintaspay/internal/modules/transactions/domain"
	"lintaspay/internal/modules/users/domain"
	wallet "lintaspay/internal/modules/wallet/domain"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.DBConfig()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&domain.User{}, &wallet.Wallet{}, &trx.Transaction{})
	c := container.NewContainer(db)

	r := router.SetupRouter(c)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
