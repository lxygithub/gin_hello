package msg

import (
	"encoding/json"
	"fmt"
	"gin_hello/kimi"
	"gin_hello/models"
	"gin_hello/openai"
	"math/rand"
	"regexp"
	"strings"

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
	//isMentioned := c.PostForm("isMentioned")
	var msgSource models.MsgSource

	json.Unmarshal([]byte(source), &msgSource)

	var replyContent string

	if msgSource.To.Payload.Name != "" {
		if strings.HasPrefix(content, "kimi") {
			replyContent = kimi.SingleChat(content, nil)
		} else {
			replyContent = openai.SingleChat(content, nil)

		}
	} else if msgSource.Room.ID != "" && strings.HasPrefix(content, "#") {
		quizz := strings.Replace(content, "#", "", 1)
		var result string
		if quizz != "" {
			if strings.HasPrefix(quizz, "kimi") {
				result = kimi.SingleChat(quizz, nil)
			} else {
				result = openai.SingleChat(quizz, nil)
			}
			replyContent = fmt.Sprintf("@%s\n %s", msgSource.From.Payload.Name, result)
		} else {
			var what = []string{"?", "？", "??", "？？", "搞咩？", "干嘛"}
			replyContent = what[rand.Intn(len(what))]
		}
	} else {
		replyContent = content
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
