package wechat

import (
	"bytes"
	"encoding/json"
	"gin_hello/models"
	"gin_hello/wechat/msg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Send_wechat_msg_api_url = "http://117.50.199.110:3001/webhook/msg/v2?token=lroRidFIwN6BXvPt5UWtPp0rROQZ3VmHRllNpQstflmaOE9G"

func Send_wechat_msg(c *gin.Context) {
	// 定义POST请求的URL

	// 准备JSON数据
	jsonData := map[string]interface{}{
		"to": c.Param("to"),
		"data": map[string]interface{}{
			"type":    "text",
			"content": c.Param("msg"),
		},
	}
	if c.Param("send_type") == "g" {
		jsonData["isRoom"] = true
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", Send_wechat_msg_api_url, bytes.NewBuffer(jsonValue))
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
func Send_wechat_msg2(c *gin.Context) {
	// 定义POST请求的URL

	// 准备JSON数据
	jsonData := map[string]interface{}{
		"to": c.PostForm("to"),
		"data": map[string]interface{}{
			"type":    "text",
			"content": c.PostForm("msg"),
		},
	}
	if c.PostForm("send_type") == "g" {
		jsonData["isRoom"] = true
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", Send_wechat_msg_api_url, bytes.NewBuffer(jsonValue))
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

func Received_wechat_msg(c *gin.Context) {
	respData := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"type":    "text",
			"content": msg.CreateReplyMsg(c),
		},
	}

	c.JSON(http.StatusOK, respData)
}