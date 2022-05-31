package middlewares

import (
	"database/sql"
	"fmt"

	"github.com/hoxito/walletgo/tools/custerror"

	"github.com/gin-gonic/gin"
	"github.com/hoxito/walletgo/tools/db"
)

/**
 * @apiDefine AuthSession
 *
 * @apiExample
 *	Session:
 *			{String} UserId
 *			{String} UserName
 *			{String} UserPassword
 *
 * @apiErrorExample 401 Unauthorized
 *    HTTP/1.1 401 Unauthorized
 *    {
 *       "error" : "Unauthorized"
 *    }
 */
// ErrorHandler a middleware to handle errors

func ValidateSession(c *gin.Context) {

	uid, err := c.Cookie("UserId")
	if err != nil {
		c.Error(custerror.Unauthorized)
		c.Abort()
	}
	uname, err := c.Cookie("UserName")
	if err != nil {
		c.Error(custerror.Unauthorized)
		c.Abort()
	}
	fmt.Println("cookies", uid, uname)

	db := db.Mysql()
	defer db.Close()
	err = db.QueryRow("SELECT name FROM user WHERE iduser = ? AND name = ? ", uid, uname).Err()

	if err != nil {
		switch err {

		case sql.ErrNoRows:
			c.Error(custerror.Unauthorized)
			c.Abort()

		default:
			panic(err)
		}
	}

}
