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
		userId,
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

func (d *Diary) CreateDiary() error {
	cmd := `insert into diaries (
		id = $1,
		userId = $2,
		fishingDate = $3, 
		place = $4, 
		diaryComment = $5, 
		rod = $6, 
		method = $7, 
		lure = $8, 
		weather = $9, 
		wind = $10,
		tide = $11,
		createdAt = $12
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err = Db.Exec(cmd,
		createUlid(),
		d.userId,
		d.fishingDate,
		d.place,
		d.diaryComment,
		d.rod,
		d.method,
		d.lure,
		d.weather,
		d.wind,
		d.tide,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (d *Diary) UpdateDiary() error {
	cmd := `update diaries set  
		userId = $1,
		fishingDate = $2, 
		place = $3, 
		diaryComment = $4, 
		rod = $5, 
		method = $6, 
		lure = $7, 
		weather = $8, 
		wind = $9,
		tide = $10,
		where id = $11
		) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err = Db.Exec(cmd,
		d.userId,
		d.fishingDate,
		d.place,
		d.diaryComment,
		d.rod,
		d.method,
		d.lure,
		d.weather,
		d.wind,
		d.tide,
	)
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
