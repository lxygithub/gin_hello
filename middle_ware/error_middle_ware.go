package middle_ware

import (
	"errors"
	"gin_hello/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type MyError struct {
	Msg  string
	Code int
}

func (e *MyError) Error() string {
	return e.Msg
}

func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//getClientIP(c.Request)
		c.Next() // 处理请求
		// 如果有错误发生
		if len(c.Errors) > 0 {
			// 这里可以根据错误类型、错误内容等进行相应处理
			for _, e := range c.Errors {
				var myError *MyError
				switch {
				case errors.As(e.Err, &myError):
					c.JSON(myError.Code, models.NewErrorResponse(myError.Code, myError.Msg))
				default:
					c.JSON(http.StatusInternalServerError,
						models.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error"))
				}
			}
			c.Abort() // 防止多次响应
		}
	}
}

// 获取客户端IP地址的函数
func getClientIP(r *http.Request) string {
	// 从请求头中获取IP地址
	clientIP := r.Header.Get("X-Forwarded-For")
	if clientIP == "" {
		// 如果没有X-Forwarded-For头，尝试从RemoteAddr获取
		clientIP = strings.Split(r.RemoteAddr, ":")[0]
	}
	return clientIP
}
