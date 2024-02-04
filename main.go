package main

import "gin_hello/db"

func main() {

	//_ = GinServer().Run(":8081") // listen and serve on 0.0.0.0:8081
	db.Connect()
}
