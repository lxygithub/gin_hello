package main

import (
	"gin_hello/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinServer() *gin.Engine {
	ginServer := gin.Default()
	ginServer.GET("/", func(c *gin.Context) {
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

	ginServer.GET("/users", func(c *gin.Context) {

		users, err := database.GetUsers()
		c.JSON(http.StatusOK, gin.H{
			"data":  users,
			"error": err,
		})
	})
	ginServer.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	return ginServer
}
