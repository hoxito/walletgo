package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Session stores UserId and UserName
//TODO. Now it just returns the user "jose"
func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("UserId", "testuser1")
	session.Set("UserName", "jose")
	session.Save()
	c.JSON(200, gin.H{"UserName": "jose"})
}
