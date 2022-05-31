package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hoxito/walletgo/models/user"
	"github.com/hoxito/walletgo/rest/middlewares"
)

// @BasePath /api/v1

// @Summary Logs a user with its username
// @Schemes
// @Description before checking if the origin wallet has funds, it sends the declared amount to the destination wallet and creates a trnsaction
// @Tags user
// @Accept json
// @Param   originId      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Param   destinationId      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Param   amount      body float64     false  "string valid"       min(1)  max(10000000)
// @Produce json
// @Success 200 {array} wallet.Wallet
// @Router /user/new [post]
func Login(c *gin.Context) {

	body := user.Login{}
	if err := c.ShouldBindJSON(&body); err != nil {

		middlewares.AbortWithError(c, err)
		return
	}
	usr, err := user.FindByLogin(&body)
	if err != nil {

		middlewares.AbortWithError(c, err)
		return
	}

	c.SetCookie("UserId", usr.UserId, 60*60*24, "/", "localhost", false, true)
	c.SetCookie("UserName", usr.Name, 60*60*24, "/", "localhost", false, true)

	c.JSON(200, usr)
}
