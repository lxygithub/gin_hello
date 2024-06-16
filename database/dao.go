package database

import (
	"gin_hello/models"
	"gin_hello/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	// 设置数据库DSN（数据源名称）
	db := GetDB()

	// 执行查询操作
	stmt, err := db.Prepare("SELECT uuid,username, nickname FROM user") // 确保这里的查询语句按实际情况修改
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var users []models.User // 切片用于存储查询结果

	rows, err := stmt.Query()
	// 遍历查询结果
	for rows.Next() {
		var u models.User // 单个用户对象
		if err := rows.Scan(&u.Uuid, &u.Username, &u.Nickname); err != nil {
			panic(err)
		}
		users = append(users, u) // 将单个用户添加到切片中
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, models.NewSuccessResponse(
		map[string]interface{}{
			"users": users,
		},
	))
}

func AddUser(username, password, nickname string) interface{} {
	// 设置数据库DSN（数据源名称）
	db := GetDB()

	// 准备插入语句
	stmt, err := db.Prepare("INSERT INTO user(username, nickname, password, register_time) VALUES (?,?,?, CURRENT_TIMESTAMP)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// 执行插入操作
	_, err = stmt.Exec(username, nickname, utils.CalculateSaltedMD5Hash(password))
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Login(username, password string) map[string]interface{} {
	db := GetDB()

	// 准备插入语句
	stmt, err := db.Prepare("SELECT uuid, nickname from user where username = ? and password = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var (
		uuid     int
		nickname string
	)
	err = stmt.QueryRow(username, utils.CalculateSaltedMD5Hash(password)).Scan(&uuid, &nickname)
	if err != nil {
		return nil
	} else {
		token, _ := utils.GenerateToken(uuid)
		return map[string]interface{}{
			"username": username,
			"nickname": nickname,
			"uuid":     uuid,
			"token":    token,
		}
	}
}
