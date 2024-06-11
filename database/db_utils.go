package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

// GetDB 获取数据库连接实例的函数
func GetDB() *sql.DB {
	if instance == nil {
		log.Fatalf("DB instance is not initialized yet")
	}
	return instance
}
