package controllers

import (
	"fmt"

	"github.com/hoxito/walletgo/models/transaction"
	"github.com/hoxito/walletgo/models/wallet"
	"github.com/hoxito/walletgo/rest/middlewares"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Sends money from a wallet to another
// @Schemes
// @Description before checking if the origin wallet has funds, it sends the declared amount to the destination wallet and creates a trnsaction
// @Tags user
// @Accept json
// @Param   originId      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Param   destinationId      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Param   amount      body float64     false  "string valid"       min(1)  max(10000000)
// @Produce json
// @Success 200 {array} wallet.Wallet
// @Router /wallet/send [post]
func SendMoney(c *gin.Context) {

	//TODO get user from db and che balance > amount
	//get
	body := transaction.TransactionRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	tr, err := transaction.CreateTransaction(&body)

	if err != nil {
		fmt.Println("erroring creating tr", err)
		middlewares.AbortWithError(c, err)
	}
	if tr == nil {
		return
	}
	c.JSON(200, tr)

}

// @BasePath /api/v1

// @Summary Gets a transaction by its id
// @Schemes
// @Description Finds in DB a transaction with the provided transactionid Parameter
// @Tags user
// @Accept json
// @Param   transactionid      path string     false  "string valid"       minlength(1)  maxlength(50)
// @Produce json
// @Success 200 {array} transaction.Transaction
// @Router /user/transactions/:transactionid [get]
func GetTransaction(c *gin.Context) {
	wallet, err := wallet.GetTransaction(c.Param("transactionid"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallet)

}

// @BasePath /api/v1

// @Summary Gets all transactions of a certain wallet
// @Schemes
// @Description Finds in DB a transaction with the provided walletid Parameter
// @Tags user
// @Accept json
// @Param   walletid      path string     false  "string valid"       minlength(1)  maxlength(50)
// @Produce json
// @Success 200 {array} transaction.Transaction
// @Router /user/wallet/:walletid/transactions [get]
func GetTransactions(c *gin.Context) {
	wallets, err := wallet.GetTransactions(c.Param("walletid"))
	if err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallets)

}
