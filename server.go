package main

import (
	"bytes"
	"encoding/json"
	"gin_hello/database"
	"gin_hello/middle_ware"
	"gin_hello/models"
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

		send_wechat_msg(c)
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

func send_wechat_msg(c *gin.Context) {
	data := map[string]interface{}{
		"to":        c.Param("to"),
		"send_type": c.Param("send_type"),
		"msg":       c.Param("msg"),
	}
	// 定义POST请求的URL
	url := "http://117.50.199.110:3001/webhook/msg/v2?token=lroRidFIwN6BXvPt5UWtPp0rROQZ3VmHRllNpQstflmaOE9G"

	// 准备JSON数据
	jsonData := map[string]interface{}{
		"to": c.Param("to"),
		"data": map[string]interface{}{
			"type":    "text",
			"content": c.Param("msg"), 
		},
	}
	if c.Param("send_type") == "g" {
		data["isRoom"] = true
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	// 设置请求头，这里是设置内容类型为JSON
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// 初始化HTTP客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理响应，例如打印状态码或读取响应体
	var bodyBytes []byte
	_, err = resp.Body.Read(bodyBytes)
	if err != nil {
		panic(err)
	}

	c.JSON(resp.StatusCode, models.NewSuccessResponse(string(bodyBytes)))
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
