package utils

import (
	"crypto/md5"
	"fmt"
	"gin_hello/config"
	"gin_hello/models"
	"github.com/golang-jwt/jwt/v4"
	"io"
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
func GenerateToken(uuid int) (string, error) {
	// 设置过期时间为 24 小时后
	exp := time.Now().Add(24 * 30 * time.Hour).Unix()
	claims := &models.CustomClaims{
		UUID: uuid,
		MapClaims: jwt.MapClaims{
			"exp": exp,
		},
	}
	// 创建一个新的令牌对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并获取完整的令牌字符串
	tokenString, err := token.SignedString([]byte(config.ReadConf("server.jwt_secret_key").(string)))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
