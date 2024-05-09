package msg

import (
	"encoding/json"
	"fmt"
	"gin_hello/models"
	"gin_hello/openai"
	"regexp"

	"github.com/gin-gonic/gin"
)

func CreateReplyMsg(c *gin.Context) string {

	/**
		{
	    # 消息类型
	    "type": "text",
	    # 消息内容
	    "content": "你好",
	    # 消息发送方的数据
	    "source": "{}",
	    # 是否被艾特
	    "isMentioned": "0",
	    # 是否自己发送给自己的消息
	    "isMsgFromSelf": "0",
	    # 被遗弃的参数
	    "isSystemEvent": "0"
	  }
	*/

	content := c.PostForm("content")
	source := c.PostForm("source")
	isMentioned := c.PostForm("isMentioned")
	var msgSource models.MsgSource

	json.Unmarshal([]byte(source), &msgSource)

	var replyContent string

	quizz := RemoveAt(content)
	if isMentioned == "1" || msgSource.To.Payload.Name != "" {
		if quizz != "" {
			result := openai.SingleChat(quizz)
			replyContent = fmt.Sprintf("@%s\n %s", msgSource.From.Payload.Name, result)
		} else {
			replyContent = fmt.Sprintf("@%s 叫我干啥？", msgSource.From.Payload.Name)
		}
	}
	return replyContent
}

func RemoveAt(content string) string {

	// 编译正则表达式来匹配 "@" 及其后的所有字符直到空格
	re := regexp.MustCompile(`@[\p{L}\p{N}\p{P}\p{Z}]* `)

	// 使用正则表达式替换匹配的部分为空字符串
	cleanedContent := re.ReplaceAllString(content, "")

	return cleanedContent
}
