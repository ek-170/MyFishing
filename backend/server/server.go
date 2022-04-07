package server

import (
	"myfishing/backend/config"
	"myfishing/backend/controllers"
	"net/http"
)

func StartMainServer() error {
	http.HandleFunc("/", controllers.Top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
