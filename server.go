package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinServer() *gin.Engine {
	ginServer := gin.Default()
	ginServer.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "测试webhook!!!!!!!!!",
		})
	})

	ginServer.GET("/user/:username/:password", func(c *gin.Context) {
		username := c.Param("username")
		password := c.Param("password")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	return ginServer
}
