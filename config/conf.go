package config

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	instance *sql.DB
	once     sync.Once
)

func ReadConf(key string) interface{} {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称没有扩展名，则需要配置此项
	viper.AddConfigPath(".")      // 查找配置文件所在的路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Fatal(err)
	}
	return viper.Get(key)
}
