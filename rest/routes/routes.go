package routes

import (
	"github.com/hoxito/walletgo/controllers"
	"github.com/hoxito/walletgo/rest/middlewares"

	"github.com/gin-gonic/gin"
)

func WalletRoute(router *gin.Engine) {
	router.POST("/v1/wallet/send", middlewares.ValidateSession, controllers.SendMoney)
	router.POST("/v1/wallet/:id", middlewares.ValidateSession, controllers.GetWallet)
	router.GET("/v1/wallet", middlewares.ValidateSession, controllers.GetWallets)

}
func LoginRoute(router *gin.Engine) {
	router.POST("/v1/login", controllers.Login)

}

func UserRoute(router *gin.Engine) {
	router.GET("/v1/user", controllers.Login)
	router.POST("/v1/user/new", controllers.CreateUser)
	router.GET("/v1/user/wallets", controllers.GetWallets)
	router.GET("/v1/user/wallet/:walletid", controllers.GetWallet)
	router.GET("/v1/user/wallet/:walletid/transactions", controllers.GetTransactions)
	router.GET("/v1/user/transactions/:transactionid", controllers.GetTransaction)
	// router.GET("/v1/user/transactions", controllers.Login)

}
