package middle_ware

import (
	"errors"
	"gin_hello/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
