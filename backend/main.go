package main

import (
	"fmt"
	"myfishing/backend/controllers"
)

func main() {
	// fmt.Println("main start")
	// fmt.Println("@@Start models.CreateTest@@")
	// cmd := `insert into test (
	// 	name,
	// 	id,
	// 	age
	// 	) values (?, ?, ?)`

	// fmt.Println("@@Db")
	// fmt.Println(models.Db)

	// fmt.Println("@@Db.Exec@@")
	// _, err := models.Db.Exec(cmd,
	// 	"direct insert",
	// 	0001,
	// 	39)
	// fmt.Println("@@Db.Exec Done@@")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	fmt.Println("start mainserver")
	controllers.StartMainServer()
}
