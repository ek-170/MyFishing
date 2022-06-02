package dto

import "myfishing/backend/models"

type TestRequest struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Age  int    `json:"age"`
}

func (trq TestRequest) ConvertTest() (t models.Test) {
	t.Name = trq.Name
	t.Id = trq.Id
	t.Age = trq.Age

	return
}

// add comment
