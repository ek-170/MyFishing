package models

import (
	"fmt"
	"log"
)

type TestRepository interface {
	GetTest(id string) (t Test, err error)
	CreateTest(t Test) (id int, err error)
	UpdateTest(t Test) (err error)
	DeleteTest(id string) (err error)
}

type Test struct {
	Name string
	Id   int
	Age  int
}

type testRepository struct {
}

func NewTestRepository() TestRepository {
	return &testRepository{}
}

// レコード参照画面で使用
func (tr *testRepository) GetTest(id string) (test Test, err error) {
	cmd := `select  name, 
		id,
		age
		from    test
		where   id = ?`
	test = Test{}

	fmt.Println("@@Start Query@@")
	err = Db.QueryRow(cmd, id).Scan(&test.Name, &test.Id, &test.Age)
	fmt.Println(test)

	return
}

func (tr *testRepository) CreateTest(t Test) (id int, err error) {
	fmt.Println("@@Start models.CreateTest@@")
	cmd := `insert into test (
		name,
		id,
		age 
		) values (?, ?, ?)`

	fmt.Println(t)

	fmt.Println("@@Db.Exec@@")
	_, err = Db.Exec(cmd,
		t.Name,
		t.Id,
		t.Age)
	fmt.Println("@@Db.Exec Done@@")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("@@Extract Created Record@@")
	err = Db.QueryRow("SELECT id FROM test ORDER BY id DESC LIMIT 1").Scan(&t.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (tr *testRepository) UpdateTest(t Test) (err error) {
	cmd := `update test set  
		name = ?,
		age  = ?
		where id = ?`

	_, err = Db.Exec(cmd,
		t.Name,
		t.Age,
		t.Id,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

func (tr *testRepository) DeleteTest(id string) (err error) {
	cmd := `delete from test where id = ?`
	_, err = Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	return
}
