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

var validPathGET_TEST = regexp.MustCompile("^/test/([A-Z0-9]+)$")
var validPathPUT_TEST = regexp.MustCompile("^/test/([A-Z0-9]+)/edit$")
var validPathDELETE_TEST = regexp.MustCompile("^/test/([A-Z0-9]+)$")

type TestHandler interface {
	HandleTest(w http.ResponseWriter, r *http.Request)
}

type testHandler struct {
	tr models.TestRepository
}

func NewTestHandler(tr models.TestRepository) TestHandler {
	return &testHandler{tr}
}

func (th *testHandler) HandleTest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTest(w, r, th)
	case http.MethodPost:
		createTest(w, r, th)
	case http.MethodPut:
		updateTest(w, r, th)
	case http.MethodDelete:
		deleteTest(w, r, th)
	default:
		http.Error(w, fmt.Sprintf("Method %s is not allowed", r.Method), 405)
	}
}

func getTest(w http.ResponseWriter, r *http.Request, th *testHandler) {
	fmt.Println("@@Start getTest()@@")
	// パスパラメータから id を抽出
	var id = ParseIdFromURL(r, validPathGET_TEST, ID_INDEX)
	fmt.Println("@@id from URL: " + id)
	if id == "" {
		http.Error(w, "Invalid ID", 404)
	}
	// GetDiaryをコール
	fmt.Println("@@Start GetTest(id)@@")
	d, err := th.tr.GetTest(id)
	if err != nil {
		http.Error(w, "Invalid ID~~~", 404)
	}
	// 取得内容をJSONにエンコード
	fmt.Println("@@Start Encode")
	err = json.NewEncoder(w).Encode(d)
	// output, _ := json.MarshalIndent(d, "", "\t\t")
	if err != nil {
		http.Error(w, "JSON ENCODING FILURE", 500)
	}
	// w にJSON書き込み
	fmt.Println("@@Start WriteHeader")
	w.Header().Set("Content-Type", "application/json")
	// w.Write(output)
}

func createTest(w http.ResponseWriter, r *http.Request, th *testHandler) {
	fmt.Println("@@createTest()@@")
	// request bodyを読み込み
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	// JSONをデコード
	// JSONを構造体にキャスト
	fmt.Println("@@read body@@")
	var trq dto.TestRequest

	fmt.Println("@@Convert JSON to Struct@@")
	json.Unmarshal(body, &trq)
	fmt.Println("trq↓")
	fmt.Println(trq)
	t := trq.ConvertTest()
	fmt.Println("t↓")
	fmt.Println(t)

	// CreateDiaryをコール
	fmt.Println("@@CreateTest(t)@@")
	id, err := th.tr.CreateTest(t)
	// wに実行結果をヘッダとともに書き込み
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// LocationにURIを書き込み
	fmt.Println("@@Set Header@@")
	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
	fmt.Println("@@Finish@@")
}

func updateTest(w http.ResponseWriter, r *http.Request, th *testHandler) {
	// request bodyを読み込み
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	// JSONをデコード
	// JSONを構造体にキャスト
	var trq dto.TestRequest
	var idStr = ParseIdFromURL(r, validPathPUT_TEST, ID_INDEX)
	if idStr == "" {
		http.Error(w, "Invalid ID", 404)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(500)
	}
	trq.Id = id

	json.Unmarshal(body, &trq)
	t := trq.ConvertTest()
	// UpdateDiaryをコール
	err = th.tr.UpdateTest(t)
	// wに実行結果をヘッダとともに書き込み
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}

func deleteTest(w http.ResponseWriter, r *http.Request, th *testHandler) {
	// パスパラメータから id を抽出
	var id = ParseIdFromURL(r, validPathDELETE_TEST, ID_INDEX)
	if id == "" {
		http.Error(w, "Invalid ID", 404)
	}
	// DeleteDiaryをコール
	err := th.tr.DeleteTest(id)
	// wに実行結果をヘッダとともに書き込み
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(204)
}
