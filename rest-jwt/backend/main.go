package main

import (
	"encoding/json"
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

// error response
type ErrResponse struct {
	Error []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Tittle string `json:"tittle"`
	Detail string `json:"detail"`
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

// gorm db migrate
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

	response := TodoResponse{1, "test", false}
	todolist := Todos{Todo: []TodoResponse{response, response}}

	jsonTodolist, err := json.Marshal(todolist)
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return
	}
	fmt.Println(string(jsonTodolist))

	errmsg := ErrorMessage{"text", "aa"}
	errResponse := ErrResponse{Error: []ErrorMessage{errmsg, errmsg, errmsg}}

	jsonResponse, err := json.Marshal(errResponse)
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return
	}
	fmt.Println(string(jsonResponse))
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
