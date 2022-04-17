package controllers

import (
	"html/template"
	"log"
	"net/http"
)

const index string = "../frontend/build/"

func Top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(index)
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func Diaries(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(index)
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
