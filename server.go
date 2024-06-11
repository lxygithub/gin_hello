package main

import (
	"gin_hello/auth"
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
	/**
	data class RecordingServiceConfig(
	    val configUrl: String = "http://117.50.199.110:8081/recording-service-config",
	    val endpoint: String = "tos-cn-beijing.volces.com",
	    val region: String = "cn-beijing",
	    val accessKey: String = "AKLTODNlZTc0N2M4MGMxNGU2ODlmZDRlYTFkYjVlZWIwYTU",
	    val secretKey: String = "T0dOaU5EUTVNV1JqWlRBeE5EbGpOR0U0TkdVelpqSTFZalprT1RSbE5qRQ==",
	    val bucketName: String = "ark-auto-created-required-2101485836-cn-beijing",
	    val maxConcurrentTaskNum: Int = 3,
	    val audioSliceSize: Int = 10 * 1024 * 1024,
	    val skipSilence: Boolean = false,
	)
	*/
	ginServer.GET("/recording-service-config", func(c *gin.Context) {

		c.JSON(http.StatusOK, models.NewSuccessResponse(
			map[string]interface{}{
				"configUrl":            "http://117.50.199.110:8081/recording-service-config",
				"endpoint":             "tos-cn-beijing.volces.com",
				"region":               "cn-beijing",
				"accessKey":            "AKLTODNlZTc0N2M4MGMxNGU2ODlmZDRlYTFkYjVlZWIwYTU",
				"secretKey":            "T0dOaU5EUTVNV1JqWlRBeE5EbGpOR0U0TkdVelpqSTFZalprT1RSbE5qRQ==",
				"bucketName":           "ark-auto-created-required-2101485836-cn-beijing",
				"maxConcurrentTaskNum": 3,
				"audioSliceSize":       10 * 1024 * 1024,
				"skipSilence":          "false",
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
