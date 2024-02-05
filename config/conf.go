package config

import (
	"github.com/spf13/viper"
	"log"
)

func ReadConf(key string) interface{} {
	viper.SetConfigName("config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")   // 如果配置文件的名称没有扩展名，则需要配置此项
	viper.AddConfigPath("../")    // 查找配置文件所在的路径

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Fatal(err)
	}
	return viper.Get(key)
}
