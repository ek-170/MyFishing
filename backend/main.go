package main

import (
	"fmt"
	"myfishing/backend/controllers"
	"myfishing/backend/models"
)

func main() {
	fmt.Println("start")
	fmt.Println(models.Db)
	fmt.Println("finish")

	controllers.StartMainServer()
}
