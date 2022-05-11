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
var validPathPOST = regexp.MustCompile("^/diaries/new$")
var validPathPUT = regexp.MustCompile("^/diaries/([A-Z0-9]{26})/edit$")
var validPathDELETE = regexp.MustCompile("^/diaries/([A-Z0-9]{26})$")

const IDPOSITION = 2

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
	// メソッド化
	// d := models.Diary{}
	// switch r.Method {
	// case http.MethodPost, http.MethodPut:
	// 	json.NewDecoder(r.Body).Decode(&d)
	// }
	switch r.Method {
	case http.MethodGet:
		getDiary(w, r, dh)
	case http.MethodPost:
		createDiary(w, r, dh)
	case http.MethodPut:
		updateDiary(w, r)
	case http.MethodDelete:
		deleteDiary(w, r)
	default:
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), 405)
	}
}

func getDiary(w http.ResponseWriter, r *http.Request, dh *diaryHandler) {
	// パスパラメータから id を抽出
	var id = ParseIdFromURL(r, validPathGET, IDPOSITION)
	// GetDiaryをコール
	d, err := dh.dr.GetDiary(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid URL"), 500)
	}
	// 取得内容をJSONにエンコード
	err = json.NewEncoder(w).Encode(d)
	// output, _ := json.MarshalIndent(d, "", "\t\t")
	if err != nil {
		http.Error(w, fmt.Sprintf("JSON ENCODING FILURE"), 500)
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
	var drq dto.DiaryRequest
	// JSONを構造体にキャスト
	json.Unmarshal(body, &drq)
	d := drq.ConvDiaryRequestToDiary()
	// CreateDiaryをコール
	id, err := dh.dr.CreateDiary(d)
	// wに実行結果をヘッダとともに書き込み
	if err != nil {
		w.WriteHeader(500)
		return
	}

	// ここから★★
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func updateDiary(w http.ResponseWriter, r *http.Request) {
	// request bodyを読み込み
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	// JSONをデコード
	// JSONを構造体にキャスト
	// CreateDiaryをコール
	// wに実行結果をヘッダとともに書き込み
}

func deleteDiary(w http.ResponseWriter, r *http.Request) {
	// パスパラメータから id を抽出
	// DeleteDiaryをコール
	// wに実行結果をヘッダとともに書き込み
}
