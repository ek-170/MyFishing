package models

import (
	"log"
	"time"
)

type DiaryRepository interface {
	GetDiary(id string) (diary Diary, err error)
	CreateDiary(d Diary) (id int, err error)
	UpdateDiary(d Diary) (err error)
	DeleteDiary(id string) (err error)
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
		caughtFish,
		diaryComment, 
		rod, 
		method, 
		lure, 
		weather, 
		wind,
		tide,
		createdAt 
		from    diaries
		where   id = ?`
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
		id,
		userId,
		fishingDate, 
		place, 
		caughtFish,
		diaryComment, 
		rod, 
		method, 
		lure, 
		weather, 
		wind,
		tide
		) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUlid(),
		d.UserId,
		nil,
		d.Place,
		d.CaughtFish,
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

	err = Db.QueryRow("SELECT id FROM diaries ORDER BY id DESC LIMIT 1").Scan(&id)

	return
}

func (dr *diaryRepository) UpdateDiary(d Diary) (err error) {
	cmd := `update diaries set  
		userId = ?,
		fishingDate = ?, 
		place = ?, 
		caughtFish = ?,
		diaryComment = ?, 
		rod = ?, 
		method = ?, 
		lure = ?, 
		weather = ?, 
		wind = ?,
		tide = ?,
		where id = ?`

	_, err = Db.Exec(cmd,
		d.UserId,
		d.FishingDate,
		d.Place,
		d.CaughtFish,
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

func (dr *diaryRepository) DeleteDiary(id string) (err error) {
	cmd := `delete from diaries where id = ?`
	_, err = Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
