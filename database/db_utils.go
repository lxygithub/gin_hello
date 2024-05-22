package database

import (
	"database/sql"
	"fmt"
	"gin_hello/auth"
	"gin_hello/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"sync"
	"time"
)

// 定义DB单例和sync.Once实例
var (
	instance *sql.DB
	once     sync.Once
)

// InitDB 初始化数据库连接的函数
func InitDB(dataSourceName string) {
	var err error
	once.Do(func() {
		instance, err = sql.Open("gin_hello", dataSourceName)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}

		// 设置连接池参数
		instance.SetMaxOpenConns(25)
		instance.SetMaxIdleConns(5)
		instance.SetConnMaxLifetime(time.Minute * 5)
	})
}

// GetDB 获取数据库连接实例的函数
func GetDB() *sql.DB {
	if instance == nil {
		log.Fatalf("DB instance is not initialized yet")
	}
	return instance
}

func GetUsers() ([]models.User, error) {
	// 设置数据库DSN（数据源名称）
	db := GetDB()
	defer db.Close()

	// 执行查询操作
	rows, err := db.Query("SELECT User FROM user") // 确保这里的查询语句按实际情况修改
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User // 切片用于存储查询结果

	// 遍历查询结果
	for rows.Next() {
		var u models.User // 单个用户对象
		if err := rows.Scan(&u.Username); err != nil {
			return nil, err
		}
		users = append(users, u) // 将单个用户添加到切片中
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// 返回用户切片
	return users, nil
}

func AddUser(username, password, nickname string) {
	// 设置数据库DSN（数据源名称）
	db := GetDB()
	defer db.Close()

	// 准备插入语句
	stmt, err := db.Prepare("INSERT INTO user(uuid, username, nickname, password, token, register_time) VALUES (?,?,?,?,?,?, CURRENT_TIMESTAMP)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// 要插入的数据
	uuid := 1000000000
	token, _ := auth.GenerateToken(strconv.Itoa(uuid))

	// 执行插入操作
	_, err = stmt.Exec(uuid, username, nickname, auth.CalculateSaltedMD5Hash(password), token)
	if err != nil {
		fmt.Println("插入数据时出错:", err)
	} else {
		fmt.Println("数据插入成功")
	}

}

func UpdateToken(uuid int) {
	db := GetDB()
	defer db.Close()

	// 准备插入语句
	stmt, err := db.Prepare("UPDATE user set token = ? where uuid = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// 要插入的数据
	token, _ := auth.GenerateToken(strconv.Itoa(uuid))

	// 执行插入操作
	_, err = stmt.Exec(uuid, token)
	if err != nil {
		fmt.Println("插入数据时出错:", err)
	} else {
		fmt.Println("数据插入成功")
	}
}

func Login(username, password string) {
	db := GetDB()
	defer db.Close()

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
	err = stmt.QueryRow(username, auth.CalculateSaltedMD5Hash(password)).Scan(&uuid, &nickname)
	if err != nil {
		fmt.Println("插入数据时出错:", err)
	} else {
	}
}
