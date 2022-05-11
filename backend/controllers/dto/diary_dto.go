package dto

import "myfishing/backend/models"

// for POST, PUT Method
type DiaryRequest struct {
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

func (drq DiaryRequest) ConvDiaryRequestToDiary() (d models.Diary) {
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

// for GET, POST, PUT Method
type DiaryResponse struct {
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

func (drs DiaryResponse) ConvDiaryResponseToDiary() (d models.Diary) {
	d.Id = drs.Id
	d.UserId = drs.UserId
	d.FishingDate = drs.FishingDate
	d.Place = drs.Place
	d.CaughtFish = drs.CaughtFish
	d.DiaryComment = drs.DiaryComment
	d.Rod = drs.Rod
	d.Method = drs.Method
	d.Lure = drs.Lure
	d.Weather = drs.Weather
	d.Wind = drs.Wind
	d.Tide = drs.Tide

	return
}
