package kimi

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

var apiUrl = "https://api.moonshot.cn/v1/chat/completions"

func SingleChat(quizz string, answerType *string) string {
	jsonBody := map[string]interface{}{
		"model": "moonshot-v1-8k",
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
		"temperature": 0.3,
	}
	reqJson, err := json.Marshal(jsonBody)
	if err != nil {
		panic(err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqJson))
	if err != nil {
		panic(err)
	}

	// 设置请求头，这里是设置内容类型为JSON
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.ReadConf("kimi.api_key").(string)))

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
		var kimiResp models.KimiResponse
		json.Unmarshal(bodyBytes, &kimiResp)
		if answerType != nil && *answerType == "complete" {
			return string(bodyBytes)
		} else {
			if len(kimiResp.Choices) > 1 {
				var answers strings.Builder
				for index, chioce := range kimiResp.Choices {
					answers.WriteString(fmt.Sprintf("回答%d: \n", index) + chioce.Message.Content + "\n")
				}
				return answers.String()
			} else {
				return kimiResp.Choices[0].Message.Content
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
