package main

import (
	"fmt"
	"myfihing/backend/models"
	"myfihing/backend/server"
)

func main() {
	fmt.Println(models.Db)

	server.StartMainServer()
}
