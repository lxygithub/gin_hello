package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// User 结构体定义用户数据的类型
type User struct {
	Username string
}

// Connect 连接数据库并返回所有用户和错误
func Connect() ([]User, error) {
	// 设置数据库DSN（数据源名称）
	dsn := "user:password@tcp(117.50.199.110:6603)/databaseName" // 注意修改这里的user, password和databaseName

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// 检查数据库连接是否成功
	err = db.Ping()
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
