package controllers

import (
	"myfishing/backend/config"
	"myfishing/backend/models"
	"net/http"
	"regexp"
)

func ParseIdFromURL(r *http.Request, validPath *regexp.Regexp, targetPath int) string {
	q := validPath.FindStringSubmatch(r.URL.Path)
	if len(q) == 0 {
		return ""
	}
	return q[targetPath]
}

func StartMainServer() error {
	var dr = models.NewDiaryRepository()
	var hd = NewDiaryHandler(dr)

	http.HandleFunc("/", Top)

	http.HandleFunc("/diaries/", hd.HandleDiary)

	// http.HandleFunc("/search", )
	// http.HandleFunc("/register", )
	// http.HandleFunc("/login", )
	// http.HandleFunc("/config", )
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
