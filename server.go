package main

import (
	"gin_hello/auth"
	"gin_hello/config"
	"gin_hello/database"
	"gin_hello/kimi"
	"gin_hello/middle_ware"
	"gin_hello/models"
	"gin_hello/openai"
	"gin_hello/wechat"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinServer() *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(middle_ware.ErrorHandlingMiddleware())
	ginServer.Use(middle_ware.AuthMiddleware())
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"message": "这是首页",
			},
		))
	})
	ginServer.POST("/send_wechat_msg/:to/:send_type/:msg", func(c *gin.Context) {
		wechat.SendWechatMsg(c)
	})
	ginServer.POST("/send_wechat_msg2", func(c *gin.Context) {
		wechat.SendWechatMsg2(c)
	})
	ginServer.POST("/received_wechat_msg", func(c *gin.Context) {
		wechat.ReceivedWechatMsg(c)
	})
	ginServer.POST("/kimi/single_chat", func(c *gin.Context) {
		kimi.Chat(c)
	})
	ginServer.POST("/chatgpt/single_chat", func(c *gin.Context) {
		openai.Chat(c)
	})

	ginServer.GET("/users", func(c *gin.Context) {
		database.GetUsers(c)

	})
	ginServer.GET("/recording-service-config", func(c *gin.Context) {

		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"configUrl":            config.ReadConf("configUrl").(string),
				"endpoint":             config.ReadConf("endpoint").(string),
				"region":               config.ReadConf("region").(string),
				"accessKey":            config.ReadConf("accessKey").(string),
				"secretKey":            config.ReadConf("secretKey").(string),
				"bucketName":           config.ReadConf("bucketName").(string),
				"maxConcurrentTaskNum": config.ReadConf("maxConcurrentTaskNum").(int),
				"audioSliceSize":       config.ReadConf("audioSliceSize").(int),
				"skipSilence":          config.ReadConf("skipSilence").(bool),
			},
		))
	})
	ginServer.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	ginServer.POST("/register", func(c *gin.Context) {
		auth.Register(c)
	})
	ginServer.POST("/login", func(c *gin.Context) {
		auth.Login(c)
	})
	return ginServer
}
