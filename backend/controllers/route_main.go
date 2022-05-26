package controllers

import (
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	//このメソッドじゃだめ？？
	// t, err := template.ParseFiles(consts.Index)
	// files := http.FileServer(http.Dir(consts.Index))
	// http.Handle("/static/", http.StripPrefix("/static/", files))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// err = t.Execute(w, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Fprintf(w, "hello\n")
}
