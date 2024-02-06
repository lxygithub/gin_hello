package main

func main() {

	//database.InitDB(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
	//	config.ReadConf("database.username").(string),
	//	config.ReadConf("database.password").(string),
	//	config.ReadConf("database.host").(string),
	//	config.ReadConf("database.port").(int),
	//	config.ReadConf("database.dbname").(string),
	//))
	//_ = GinServer().Run(fmt.Sprintf(":%d", config.ReadConf("server.port").(int)))
	_ = GinServer().Run(":8081")

}
