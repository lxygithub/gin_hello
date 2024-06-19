package middle_ware

import (
	"fmt"
	"gin_hello/config"
	"gin_hello/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

var noAuth = []string{
	"/",
	"/login",
	"/register",
	"/recording-service-config",
}

func noNeedAuth(noAuth []string, path string) bool {
	for _, i := range noAuth {
		if i == path {
			return true
		}
	}
	return false
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if noNeedAuth(noAuth, c.Request.URL.Path) {
			c.Next()
		} else {
			// 从请求中获取token，通常是从Authorization头部或请求参数中获取
			tokenString := c.GetHeader("auth_token")

			// 解析和验证token
			token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{},
				func(token *jwt.Token) (interface{}, error) {
					// 提供用于验证签名的密钥
					return []byte(config.ReadConf("server.jwt_secret_key").(string)), nil
				})

			// 如果token无效或解析时出错，则拒绝访问
			if err != nil || !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized,
					models.NewErrorResponse(http.StatusUnauthorized, "Invalid token"))
				return
			}
			// token有效，继续处理请求
			if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
				fmt.Print(claims.UUID)
			}
			c.Next()
		}
	}
}
