package controllers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/hoxito/walletgo/models/user"
	"github.com/hoxito/walletgo/rest/middlewares"

	"github.com/gin-gonic/gin"
)

/**
 * @apiDefine CreateUser
 *
 * @apiExample
 * 	 Name     string required
 *   Email    string email
 *   Password string required
 *
 */
func CreateUser(c *gin.Context) {

	//TODO get user from db and che balance > amount
	//get
	body := user.UserRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	tr, err := user.CreateUser(&body)

	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	if tr == nil {
		return
	}
	c.JSON(200, tr)

}

//Gets wallet from DB. TODO
func CreateWallet(c *gin.Context) {
	c.Param("walletId")

	c.JSON(200, gin.H{"WalletId": c.Param("walletId")})

}

//Gets a wallet from DB of a certain user
func GetWallet(c *gin.Context) {
	session := sessions.Default(c)
	uid := fmt.Sprint(session.Get("UserId"))
	wallet, err := user.GetWallet(uid, c.Param("walletid"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallet)

}

//Gets wallets of a certain user
func GetWallets(c *gin.Context) {
	session := sessions.Default(c)
	uid := fmt.Sprint(session.Get("UserId"))
	wallets, err := user.GetWallets(uid)
	if err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallets)

}
