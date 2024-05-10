package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin_hello/config"
	"gin_hello/models"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var ApiUrl = "https://cfcus02.opapi.win/v1/chat/completions"

func SingleChat(quizz string, answerType *string) string {
	SendMsg(quizz)
	jsonBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": config.ReadConf("chatgpt.prompt").(string),
			},
			{
				"role":    "user",
				"content": quizz,
			},
		},
	}
	reqJson, err := json.Marshal(jsonBody)
	if err != nil {
		panic(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", ApiUrl, bytes.NewBuffer(reqJson))
	if err != nil {
		panic(err)
	}

	// 设置请求头，这里是设置内容类型为JSON
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.ReadConf("chatgpt.api_key").(string)))

	// 初始化HTTP客户端
	client := &http.Client{}

	client.Timeout = 30 * time.Second
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		// 处理响应，例如打印状态码或读取响应体
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var chatgptResp models.ChatgptResp
		json.Unmarshal(bodyBytes, &chatgptResp)
		if answerType != nil && *answerType == "complete" {
			return string(bodyBytes)
		} else {
			if len(chatgptResp.Choices) > 1 {
				var answers strings.Builder
				for index, chioce := range chatgptResp.Choices {
					answers.WriteString(fmt.Sprintf("回答%d: \n", index) + chioce.Message.Content + "\n")
				}
				return answers.String()
			} else {
				return chatgptResp.Choices[0].Message.Content
			}
		}
	} else {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var kimiErr models.KimiErrResp
		json.Unmarshal(bodyBytes, &kimiErr)

		return kimiErr.Error.Message
	}
}

func Chat(c *gin.Context) {
	quizz := c.PostForm("quizz")
	answerType := c.PostForm("answerType")
	result := SingleChat(quizz, &answerType)
	c.JSON(http.StatusOK, models.NewSuccessResponse(result))
}
func SendMsg(quizz string) {
	// 定义POST请求的URL

	// 准备JSON数据
	jsonData := map[string]interface{}{
		"to": "相见不如怀念",
		"data": map[string]interface{}{
			"type":    "text",
			"content": quizz,
		},
	}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", config.ReadConf("wechat.api_url").(string), bytes.NewBuffer(jsonValue))
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
}
