package models

import "time"

type Diary struct {
	id           string
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
	created_at   time.Time
}
