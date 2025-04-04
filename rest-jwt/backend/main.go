package main

import (
	"fmt"
)

// response
type Todos struct {
	Todo []TodoResponse `json:"todos"`
}
type TodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ErrResponse struct {
	Message map[string]any `json:"message"`
}

type TodoMethod interface {
	CheckTitle() error
}

// request
type TodoCreateRequest struct {
	Title string `json:"title"`
}

func (t TodoCreateRequest) CheckTitle() error {
	if t.Title == "" || len(t.Title) > 255 {
		return fmt.Errorf("Todo Titleがないか、255文字以上です")
	}
	return nil
}

type TodoUpdateRequest struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t TodoUpdateRequest) CheckTitle() error {
	if t.Title == "" || len(t.Title) > 255 {
		return fmt.Errorf("Todo Titleがないか、255文字以上です")
	}
	return nil
}

func (t TodoUpdateRequest) CheckId() error {
	if t.Id <= 0 {
		return fmt.Errorf("Todo IDがないか、0以下です")
	}
	return nil
}

// gorm db mygrate
type TodoTable struct {
	ID        int    `gorm:"primary_key;auto_increment"`
	Title     string `gorm:"size:255"`
	Completed bool   `gorm:"default:false"`
}

func main() {
	var CreateRequest TodoMethod = &TodoCreateRequest{""}
	// var UpdateRequest TodoMethod = &TodoUpdateRequest{1, "a", false}
	UpdateRequest := TodoUpdateRequest{0, "", false}
	if err := CreateRequest.CheckTitle(); err != nil {
		fmt.Println(err)
	}
	if err := UpdateRequest.CheckTitle(); err != nil {
		fmt.Println(err)
	}
	if err := UpdateRequest.CheckId(); err != nil {
		fmt.Println(err)
	}
}

/*
Todo
DB-struct
gorm
gin
JWT
api-end-point{
	/login[
		POST
	]

	/signup[
		POST
	]

	/todo[
		GET
		POST
		PUT
		DELETE
	]
}

*/
