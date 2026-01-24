package router

import (
	"lintaspay/internal/app/container"
	"lintaspay/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(c *container.Container) *gin.Engine {
	r := gin.Default()

	// USER / AUTH ROUTES
	auth := r.Group("/auth")
	{
		auth.POST("/register", c.User.AuthHandler.Register)
		auth.POST("/login", c.User.AuthHandler.Login)

	}
	wallet := r.Group("/wallet")
	wallet.Use(middleware.AuthMiddleware())

	{
		wallet.POST("/create", c.Wallet.WalletHandler.Create)

	}
	transactions := r.Group("/trx")
	transactions.Use(middleware.AuthMiddleware())

	{
		transactions.POST("/top-up", c.Trx.TrxHandler.TopUp)
		transactions.POST("/transfer", c.Trx.TrxHandler.Transfer)

	}

	return r
}
