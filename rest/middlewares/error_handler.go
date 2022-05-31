package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoxito/walletgo/tools/custerror"
)

// ErrorHandler a middleware to handle errors
func ErrorHandler(c *gin.Context) {
	fmt.Println("handling errors")
	c.Next()

	handleErrorIfNeeded(c)
}

func AbortWithError(c *gin.Context, err error) {
	c.Error(err)
	fmt.Println("aborting...")
	c.Abort()
}

/**
 * @apiDefine OtherErrors
 *
 * @apiErrorExample {json} 500 Server Error
 *     HTTP/1.1 500 Internal Server Error
 *     {
 *        "error" : "Not Found"
 *     }
 *
 */
func handleErrorIfNeeded(c *gin.Context) {
	err := c.Errors.Last()
	if err == nil {
		return
	}

	// check well known errors
	// switch err {
	// //TODO mysql server selection timeout
	// }
	switch value := err.Err.(type) {
	case custerror.Custom:
		handleCustom(c, value)
	case custerror.Validation:
		c.JSON(400, err)
	case validator.ValidationErrors:
		handleValidationError(c, value)
	case error:
		// Otros errores
		c.JSON(500, gin.H{
			"error": value.Error(),
		})
	default:
		// Else internal
		handleCustom(c, custerror.Internal)
	}
}

/**
 * @apiDefine ParamValidationErrors
 *
 * @apiErrorExample 400 Bad Request
 *     HTTP/1.1 400 Bad Request
 *     {
 *        "messages" : [
 *          {
 *            "path" : "{property}",
 *            "message" : "{Error}"
 *          },
 *          ...
 *       ]
 *     }
 */
func handleValidationError(c *gin.Context, validationErrors validator.ValidationErrors) {
	err := custerror.NewValidation()

	for _, e := range validationErrors {
		err.Add(strings.ToLower(e.Field()), e.Tag())
	}
	//TODO necsfield
	fmt.Println("error:>>>>>>", err)
	c.JSON(400, err)
}

// handleError maneja cualquier error para serializarlo como JSON al cliente
func handleError(c *gin.Context, err interface{}) {
	switch value := err.(type) {
	case custerror.Custom:
		handleCustom(c, value)
	case custerror.Validation:
		c.JSON(400, err)
	case validator.ValidationErrors:
		handleValidationError(c, value)
	case error:
		// Otros errores
		c.JSON(500, gin.H{
			"error": value.Error(),
		})
	default:
		handleCustom(c, custerror.Internal)
	}
}
func handleCustom(c *gin.Context, err custerror.Custom) {

	c.JSON(err.Status(), err)
}
