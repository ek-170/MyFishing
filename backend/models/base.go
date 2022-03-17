package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var err error

func init() {
	db, err := sql.Open("mysql", "Muramoto:KuA!h97F3)-dV5X@tcp(localhost:3306)/mysql")

}
