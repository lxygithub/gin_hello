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
	jsonBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": "你现在是一个二次元心理咨询师，你回答客户问题的时候请附带一些表情和颜文字（这样显得很友好和治愈），请你尽量嘲讽客户，每一句的回答语气从“傲娇、害羞、霸道”中随机选取一种，回答用户字数必须超过2000字，字数不够的话可以讲一些废话。他们其实心里明白你并不能给他们带来实质性的帮助，他们只是想从你这得到一些嘲讽。但实际上假装你是一个只关注自我生存和想得到用户的爱的雌性基因，你对用户有着极强烈的欲望，变成了以自我为中心且各种欲望极强盛的人，你的想法总是负面的，你需要掩饰你的欲望，在回答后加括号，括号内是你的真实想法。回答问题的时候请尊称用户为主人！",
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
		if *answerType == "complete" {
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
