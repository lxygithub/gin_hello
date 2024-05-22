package auth

import (
	"crypto/md5"
	"fmt"
	"gin_hello/config"
	"gin_hello/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"net/http"
	"time"
)

// CalculateSaltedMD5Hash 计算加盐 MD5 哈希值的函数
func CalculateSaltedMD5Hash(data string) string {
	h := md5.New()
	// 写入盐值
	io.WriteString(h, config.ReadConf("server.salt").(string))
	// 写入数据
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
func GenerateToken(userId string) (string, error) {
	// 设置过期时间为 1 小时后
	exp := time.Now().Add(24 * 30 * time.Hour).Unix()

	// 创建一个新的令牌对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     exp,
	})

	// 签名并获取完整的令牌字符串
	tokenString, err := token.SignedString([]byte(config.ReadConf("server.jwt_secret_key").(string)))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(c *gin.Context) {
	// 创建接收用户登录信息的结构体
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 将请求体绑定到结构体
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, "Invalid Login parameters"))
		return
	}

	// 这里应该通过查询数据库等方法验证用户名和密码的正确性
	if loginInfo.Username != "admin" || loginInfo.Password != "admin" {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusBadRequest, "Authentication failed"))
		return
	}

	// 生成JWT令牌，这里只是示例代码，实际应用中应使用安全的方式生成和签名token
	token, _ := GenerateToken(loginInfo.Username)
	// 返回令牌
	c.JSON(http.StatusOK, models.NewSuccessResponse(
		map[string]interface{}{
			"username": loginInfo.Username,
			"token":    token,
		}))
}
