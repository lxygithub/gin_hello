package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() {
	// 设置数据库DSN（数据源名称）
	dsn := "mysql:password@tcp(117.50.199.110:6603)/mysql"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查数据库连接是否成功
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 这里可以执行查询操作
	// 例如，查询所有用户
	rows, err := db.Query("SELECT User FROM user")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 遍历查询结果
	for rows.Next() {
		var (
			User string
		)
		if err := rows.Scan(&User); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User: %s\n", User)
	}

	// 检查遍历过程中是否有错误发生
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 其他数据库操作...
}
