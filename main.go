package main

import (
	"fmt"
	"gin_hello/config"
	"gin_hello/database"
	"log"
)

func main() {
	mysql := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.ReadConf("database.username").(string),
		config.ReadConf("database.password").(string),
		config.ReadConf("database.host").(string),
		config.ReadConf("database.port").(int),
		config.ReadConf("database.dbname").(string),
	)
	log.Println(mysql)
	database.InitDB(mysql)
	_ = GinServer().RunTLS(fmt.Sprintf(":%d", config.ReadConf("server.port").(int)), "/etc/ssl/certs/cloudflare-origin.crt", "/etc/ssl/private/cloudflare-origin.key")

}
