package models

import (
	"log"
	"time"
)

type Diary struct {
	id           string
	userId       string
	fishingDate  string
	place        string
	caughtFish   string
	diaryComment string
	rod          string
	method       string
	lure         string
	weather      string
	wind         int
	tide         string
	createdAt    time.Time
}

func GetDiary(id int) (diary Diary, err error) {
	cmd := `select  id, 
		fishingDate, 
		place, 
		diaryComment, 
		rod, 
		method, 
		lure, 
		weather, 
		wind,
		tide,
		createdAt 
		from    diaries
		where   id = $1`
	diary = Diary{}

	err = Db.QueryRow(cmd, id).Scan(
		&diary.id,
		&diary.fishingDate,
		&diary.place,
		&diary.caughtFish,
		&diary.diaryComment,
		&diary.rod,
		&diary.method,
		&diary.lure,
		&diary.weather,
		&diary.wind,
		&diary.tide,
		&diary.createdAt)

	return diary, err
}

func (u *User) CreateDiary(  ) (err error) {
	cmd := `insert into diaries (
		) values (  )`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (d *Diary) UpdateDiary() error {
	cmd := `update diaries set  更新項目を書く
	where id = $2`
	_, err = Db.Exec(cmd, ,d.id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (d *Diary) DeleteDiary() error {
	cmd := `delete from diaries where id = $1`
	_, err = Db.Exec(cmd, d.id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
