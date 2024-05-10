package database

import (
	"database/sql"
	"log"
	"sync"
	"time"
)

// 定义DB单例和sync.Once实例
var (
	instance *sql.DB
	once     sync.Once
)

// 初始化数据库连接的函数
func InitDB(dataSourceName string) {
	var err error
	once.Do(func() {
		instance, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatalf("Error opening database: %v", err)
		}

		// 设置连接池参数
		instance.SetMaxOpenConns(25)
		instance.SetMaxIdleConns(5)
		instance.SetConnMaxLifetime(time.Minute * 5)
	})
}

// 获取数据库连接实例的函数
func GetDB() *sql.DB {
	if instance == nil {
		log.Fatalf("DB instance is not initialized yet")
	}
	return instance
}

// User 结构体定义用户数据的类型
type User struct {
	Username string
}

func GetUsers() ([]User, error) {
	// 设置数据库DSN（数据源名称）
	db := GetDB()
	defer db.Close()

	// 检查数据库连接是否成功
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	// 执行查询操作
	rows, err := db.Query("SELECT User FROM user") // 确保这里的查询语句按实际情况修改
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User // 切片用于存储查询结果

	// 遍历查询结果
	for rows.Next() {
		var u User // 单个用户对象
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
