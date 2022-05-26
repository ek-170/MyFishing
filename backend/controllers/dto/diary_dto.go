package dto

import "myfishing/backend/models"

type DiaryRequest struct {
	Id           string `json:"id"`
	UserId       string `json:"userId"`
	FishingDate  string `json:"date"`
	Place        string `json:"place"`
	CaughtFish   string `json:"caughtFish"`
	DiaryComment string `json:"comment"`
	Rod          string `json:"rod"`
	Method       string `json:"method"`
	Lure         string `json:"lure"`
	Weather      string `json:"weather"`
	Wind         int    `json:"wind"`
	Tide         string `json:"tide"`
}

func (drq DiaryRequest) ConvertDiary() (d models.Diary) {
	d.UserId = drq.UserId
	d.FishingDate = drq.FishingDate
	d.Place = drq.Place
	d.CaughtFish = drq.CaughtFish
	d.DiaryComment = drq.DiaryComment
	d.Rod = drq.Rod
	d.Method = drq.Method
	d.Lure = drq.Lure
	d.Weather = drq.Weather
	d.Wind = drq.Wind
	d.Tide = drq.Tide

	return
}
