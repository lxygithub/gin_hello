package auth

import (
	"gin_hello/database"
	"gin_hello/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

	userInfo := database.Login(loginInfo.Username, loginInfo.Password)

	if userInfo != nil {
		c.JSON(http.StatusOK, models.NewSuccessResponse(userInfo))
	} else {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, "该用户不存在"))
	}
}
func Register(c *gin.Context) {
	// 创建接收用户登录信息的结构体
	var RegisterInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}

	// 将请求体绑定到结构体
	if err := c.ShouldBindJSON(&RegisterInfo); err != nil {
		c.JSON(http.StatusBadRequest, models.NewErrorResponse(http.StatusBadRequest, "Invalid Register parameters"))
		return
	}

	err := database.AddUser(RegisterInfo.Username, RegisterInfo.Password, RegisterInfo.Nickname)

	if err != nil {
		c.JSON(http.StatusUnauthorized, models.NewErrorResponse(http.StatusUnauthorized, "注册失败"))
	} else {
		c.JSON(http.StatusOK, models.NewSuccessResponse("注册成功"))
	}
}
