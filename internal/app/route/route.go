package router

import (
	"lintaspay/internal/app/container"
	"lintaspay/internal/app/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(c *container.Container) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		transactions.GET("/history", c.Trx.TrxHandler.HistoryTransaction)

	}

	return r
}
