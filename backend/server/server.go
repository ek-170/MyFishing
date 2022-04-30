package server

import (
	"myfishing/backend/config"
	"myfishing/backend/controllers"
	"net/http"
	"regexp"
	"strconv"
)

func StartMainServer() error {
	// http.Handle("/", http.FileServer(http.Dir("../frontend/build/")))
	http.HandleFunc("/", controllers.Top)
	http.HandleFunc("/diaries", controllers.Diaries)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}

// ここから修正↓
var validPath = regexp.MustCompile("^/diaries/(edit|update|delete)/([0-9]+)$")

func ParseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		id, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, id)
	}
}
