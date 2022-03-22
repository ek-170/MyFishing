package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oklog/ulid/v2"
)

var Db *sql.DB

var err error

func init() {
	Db, err := sql.Open("mysql", "Muramoto:KuA!h97F3)-dV5X@tcp(localhost:3306)/mysql")
	if err != nil {
		log.Fatalln(err)
	}
}

func createULID() string {
	t := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return fmt.Sprint(ulid.MustNew(ulid.Timestamp(t), entropy))
}

func Encrypt(plaintext string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
}
