package main

import (
	"gin_hello/database"
	"gin_hello/kimi"
	"gin_hello/middle_ware"
	"gin_hello/models"
	"gin_hello/openai"
	"gin_hello/wechat"
	"net/http"

	"github.com/gin-gonic/gin"
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

	ginServer.POST("/login", func(c *gin.Context) {
		login(c)
	})
	return ginServer
}

func login(c *gin.Context) {
	// 创建接收用户登录信息的结构体
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 将请求体绑定到结构体
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, "Invalid login parameters"))
		return
	}

	// 这里应该通过查询数据库等方法验证用户名和密码的正确性
	if loginInfo.Username != "admin" || loginInfo.Password != "admin" {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusBadRequest, "Authentication failed"))
		return
	}

	// 生成JWT令牌，这里只是示例代码，实际应用中应使用安全的方式生成和签名token
	token := "some.jwt.token"
	// 返回令牌
	c.JSON(http.StatusOK, models.NewSuccessResponse(
		map[string]interface{}{
			"token": token,
		}))
}
