package main

import (
	"fmt"
	"myfishing/backend/models"
	"myfishing/backend/server"
)

func main() {
	fmt.Println("start")
	fmt.Println(models.Db)
	fmt.Println("finish")

	server.StartMainServer()
}
