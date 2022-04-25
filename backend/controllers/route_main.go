package controllers

import (
	"html/template"
	"log"
	"myfishing/backend/consts"
	"myfishing/backend/models"
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(consts.Index)
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func Diaries(w http.ResponseWriter, r *http.Request) {
	d := models.Diary{}
	switch r.Method {
	case "POST", "PUT":
		d = models.Diary{
			// Diaryの値を埋める
		}
	}
	switch r.Method {
	case "GET":
		_, err := models.GetDiary(i)
		if err != nil {
			log.Fatalln(err)
		}
	case "POST":
		d.CreateDiary()
	case "PUT":
		d.UpdateDiary()
	case "DELETE":
		models.DeleteDiary()
		// defaultでHTTP 405 Error
	}
}
