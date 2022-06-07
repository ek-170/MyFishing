package main

import (
	"fmt"
	"myfishing/backend/controllers"
)

func main() {
	fmt.Println("start mainserver")
	controllers.StartMainServer()
}
