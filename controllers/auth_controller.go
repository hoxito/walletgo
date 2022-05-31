package controllers

import (
	"github.com/gin-gonic/gin"
)

//Session stores UserId and UserName
//TODO. Now it just returns the user "jose"
func Login(c *gin.Context) {
	c.SetCookie("UserId", "testuser1", 60*60*24, "/", "localhost", false, true)
	c.SetCookie("UserName", "jose", 60*60*24, "/", "localhost", false, true)
	c.JSON(200, gin.H{"UserName": "jose"})
}
