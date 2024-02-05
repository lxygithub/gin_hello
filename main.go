package main

import (
	"fmt"
	"gin_hello/config"
	"gin_hello/database"
)

func main() {

	database.InitDB(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.ReadConf("database.username"),
		config.ReadConf("database.password"),
		config.ReadConf("database.host"),
		config.ReadConf("database.port"),
		config.ReadConf("database.dbname"),
	))
	_ = GinServer().Run(fmt.Sprintf(":%d", config.ReadConf("server.port")))

}
