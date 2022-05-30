package controllers

import (
	"fmt"

	"github.com/hoxito/walletgo/models/transaction"
	"github.com/hoxito/walletgo/models/wallet"
	"github.com/hoxito/walletgo/rest/middlewares"

	"github.com/gin-gonic/gin"
)

/**
 * @apiDefine SendMoney
 *
 * @apiExample
 *			{string} OriginId
 *			{string}	DestinationId
 *   		{float64}	Amount
 *
 * @apiErrorExample 400
 *    HTTP/1.1 400 Bad Request
 *    {
 *       "error" : "Not Enough Ballance"
 *    }
 */
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

//Gets a transaction from DB of a certain wallet
func GetTransaction(c *gin.Context) {
	wallet, err := wallet.GetTransaction(c.Param("transactionid"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallet)

}

//Gets transactions of a certain wallet
func GetTransactions(c *gin.Context) {
	wallets, err := wallet.GetTransactions(c.Param("walletid"))
	if err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallets)

}
