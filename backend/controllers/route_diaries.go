package controllers

import (
	"encoding/json"
	"fmt"
	"myfishing/backend/controllers/dto"
	"myfishing/backend/models"
	"net/http"
	"regexp"
	"strconv"
)

var validPathGET = regexp.MustCompile("^/diaries/([A-Z0-9]{26})$")
var validPathPUT = regexp.MustCompile("^/diaries/([A-Z0-9]{26})/edit$")
var validPathDELETE = regexp.MustCompile("^/diaries/([A-Z0-9]{26})$")

const ID_INDEX = 1

type DiaryHandler interface {
	HandleDiary(w http.ResponseWriter, r *http.Request)
}

type diaryHandler struct {
	dr models.DiaryRepository
}

func NewDiaryHandler(dr models.DiaryRepository) DiaryHandler {
	return &diaryHandler{dr}
}

func (dh *diaryHandler) HandleDiary(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDiary(w, r, dh)
	case http.MethodPost:
		createDiary(w, r, dh)
	case http.MethodPut:
		updateDiary(w, r, dh)
	case http.MethodDelete:
		deleteDiary(w, r, dh)
	default:
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), 405)
	}
}

func getDiary(w http.ResponseWriter, r *http.Request, dh *diaryHandler) {
	// パスパラメータから id を抽出
	var id = ParseIdFromURL(r, validPathGET, ID_INDEX)
	if id == "" {
		http.Error(w, "Invalid ID", 404)
	}
	// GetDiaryをコール
	d, err := dh.dr.GetDiary(id)
	if err != nil {
		http.Error(w, "Invalid ID", 404)
	}
	// 取得内容をJSONにエンコード
	err = json.NewEncoder(w).Encode(d)
	// output, _ := json.MarshalIndent(d, "", "\t\t")
	if err != nil {
		http.Error(w, "JSON ENCODING FAILURE", 500)
	}
	// w にJSON書き込み
	w.Header().Set("Content-Type", "application/json")
	// w.Write(output)
}

func createDiary(w http.ResponseWriter, r *http.Request, dh *diaryHandler) {
	// request bodyを読み込み
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	// JSONをデコード
	// JSONを構造体にキャスト
	var drq dto.DiaryRequest
	json.Unmarshal(body, &drq)
	d := drq.ConvertDiary()
	// CreateDiaryをコール
	fmt.Println("@d")
	fmt.Println(d)
	id, err := dh.dr.CreateDiary(d)
	// wに実行結果をヘッダとともに書き込み
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// LocationにURIを書き込み
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func updateDiary(w http.ResponseWriter, r *http.Request, dh *diaryHandler) {
	// request bodyを読み込み
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	// JSONをデコード
	// JSONを構造体にキャスト
	var drq dto.DiaryRequest
	var id = ParseIdFromURL(r, validPathPUT, ID_INDEX)
	if id == "" {
		http.Error(w, "Invalid ID", 404)
	}
	drq.Id = id
	json.Unmarshal(body, &drq)
	d := drq.ConvertDiary()
	// UpdateDiaryをコール
	err := dh.dr.UpdateDiary(d)
	// wに実行結果をヘッダとともに書き込み
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func deleteDiary(w http.ResponseWriter, r *http.Request, dh *diaryHandler) {
	// パスパラメータから id を抽出
	var id = ParseIdFromURL(r, validPathDELETE, ID_INDEX)
	if id == "" {
		http.Error(w, "Invalid ID", 404)
	}
	// DeleteDiaryをコール
	err := dh.dr.DeleteDiary(id)
	// wに実行結果をヘッダとともに書き込み
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}
