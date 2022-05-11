package models

import (
	"log"
	"time"
)

type User struct {
	id        int
	name      string
	email     string
	password  string
	createdAt time.Time
	diarys    []Diary
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		id,
		name,
		email,
		password,
		created_at) values ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd,
		createUlid(),
		u.name,
		u.email,
		Encrypt(u.password),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}
