package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hoxito/walletgo/models/user"
	"github.com/hoxito/walletgo/models/wallet"
	"github.com/hoxito/walletgo/rest/middlewares"
)

// @BasePath /api/v1

// @Summary Creates a user
// @Schemes
// @Description Creates a user with given name, mail and (not implemented)password
// @Tags user
// @Accept json
// @Param   name      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Param   email      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Param   password      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Produce json
// @Success 200 {array} wallet.Wallet
// @Router /user/new [post]
func CreateUser(c *gin.Context) {

	//get
	body := user.UserRequest{}
	if err := c.ShouldBindJSON(&body); err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	usr, err := user.CreateUser(&body)

	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	if usr == nil {
		return
	}
	c.JSON(200, usr)

}

// @BasePath /api/v1

// @Summary Gets a user´s current total ballance
// @Schemes
// @Description Sums every user wallets ballances for the given Currency
// @Tags user
// @Accept json
// @Param   currency      path string     false  "string valid"       minlength(1)  maxlength(50)
// @Produce json
// @Success 200 {array} wallet.BallanceResponse
// @Router /user/Ballance/:Currency [get]
func GetTotalBalance(c *gin.Context) {
	uid, err := c.Cookie("UserId")
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	wallets, err := user.Wallets(uid)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	err = validate.Var(c.Param("Currency"), "required,max=3")
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	Ballance := wallet.SumBallances(wallets, c.Param("Currency"))
	c.JSON(200, Ballance)

}

//Creates a wallet  TODO
func CreateWallet(c *gin.Context) {
	c.Param("walletId")

	c.JSON(200, gin.H{"WalletId": c.Param("walletId")})

}

// @BasePath /api/v1

// @Summary Gets a user´s wallet with the wallet´s Id
// @Schemes
// @Description Finds in DB a wallet with the provided walletid Parameter
// @Tags user
// @Accept json
// @Param   walletid      path string     false  "string valid"       minlength(1)  maxlength(50)
// @Produce json
// @Success 200 {array} wallet.Wallet
// @Router /user/wallet/:walletid [get]
func GetWallet(c *gin.Context) {

	fmt.Println("uid", c)
	uid, err := c.Cookie("UserId")
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	wallet, err := user.Wallet(uid, c.Param("walletid"))
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallet)

}

// @BasePath /api/v1

// @Summary Gets the logged user
// @Schemes
// @Description Finds in DB a user with the userId Parameter provided by cookies
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} user.User
// @Router /user [get]
func GetCurrenUser(c *gin.Context) {

	fmt.Println("uid", c)
	uid, err := c.Cookie("UserId")
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	user, err := user.Get(uid)
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, user)

}

// @BasePath /api/v1

// @Summary Gets the user´s wallets with their corresponding data
// @Schemes
// @Description Finds every wallet in DB for the logged user using session parameters
// @Tags user
// @Accept json
// @Param   walletid      path string     false  "string valid"       minlength(1)  maxlength(50)
// @Produce json
// @Success 200 {array} []wallet.Wallet
// @Router /user/wallets [get]
func GetWallets(c *gin.Context) {
	uid, err := c.Cookie("UserId")
	if err != nil {
		middlewares.AbortWithError(c, err)
		return
	}
	wallets, err := user.Wallets(uid)
	if err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	c.JSON(200, wallets)

}
