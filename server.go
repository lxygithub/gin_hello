package main

import (
	"gin_hello/database"
	"gin_hello/middle_ware"
	"gin_hello/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinServer() *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(middle_ware.ErrorHandlingMiddleware())
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"message": "这是首页",
			},
		))
	})

	ginServer.GET("/user/:username/:password", func(c *gin.Context) {
		data := map[string]interface{}{
			"username": c.Param("username"),
			"password": c.Param("password"),
		}
		c.JSON(http.StatusOK, models.NewSuccessResponse(data))
	})

	ginServer.GET("/users", func(c *gin.Context) {
		users, _ := database.GetUsers()
		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"users": users,
			},
		))
	})
	ginServer.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	return ginServer
}
