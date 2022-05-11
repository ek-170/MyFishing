package controllers

import (
	"html/template"
	"log"
	"myfishing/backend/consts"
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
