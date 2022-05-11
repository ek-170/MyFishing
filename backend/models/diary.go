package models

import (
	"log"
	"time"
)

type DiaryRepository interface {
	GetDiary(id string) (diary Diary, err error)
	CreateDiary(d Diary) (id int, err error)
	UpdateDiary(d Diary) (err error)
	DeleteDiary(d Diary) (err error)
}

type Diary struct {
	Id           string
	UserId       string
	FishingDate  string
	Place        string
	CaughtFish   string
	DiaryComment string
	Rod          string
	Method       string
	Lure         string
	Weather      string
	Wind         int
	Tide         string
	CreatedAt    time.Time
}

type diaryRepository struct {
}

func NewDiaryRepository() DiaryRepository {
	return &diaryRepository{}
}

// レコード参照画面で使用
func (dr *diaryRepository) GetDiary(id string) (diary Diary, err error) {
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
		&diary.Id,
		&diary.FishingDate,
		&diary.Place,
		&diary.CaughtFish,
		&diary.DiaryComment,
		&diary.Rod,
		&diary.Method,
		&diary.Lure,
		&diary.Weather,
		&diary.Wind,
		&diary.Tide,
		&diary.CreatedAt)

	return
}

func (dr *diaryRepository) CreateDiary(d Diary) (id int, err error) {
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
		d.UserId,
		d.FishingDate,
		d.Place,
		d.DiaryComment,
		d.Rod,
		d.Method,
		d.Lure,
		d.Weather,
		d.Wind,
		d.Tide,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	err = Db.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1").Scan(&id)

	return
}

func (dr *diaryRepository) UpdateDiary(d Diary) (err error) {
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
		d.UserId,
		d.FishingDate,
		d.Place,
		d.DiaryComment,
		d.Rod,
		d.Method,
		d.Lure,
		d.Weather,
		d.Wind,
		d.Tide,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (dr *diaryRepository) DeleteDiary(d Diary) (err error) {
	cmd := `delete from diaries where id = $1`
	_, err = Db.Exec(cmd, d.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
