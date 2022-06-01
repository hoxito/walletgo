package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hoxito/walletgo/models/user"
	"github.com/hoxito/walletgo/rest/middlewares"
)

// @BasePath /api/v1

// @Summary Logs a user with its username and password
// @Schemes
// @Description Logs the user by running a sql query against the DB checking wether if a user with the provided username and password exists. If it exists, returns it and stores User Id in session
// @Tags user
// @Accept json
// @Param   name      body string     false  "string valid"       minlength(1)  maxlength(50)
// @Param   password       body string     false  "string valid"       minlength(1)  maxlength(50)
// @Produce json
// @Success 200 {array} wallet.Wallet
// @Router /login [post]
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
