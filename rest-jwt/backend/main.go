package main

import (
	"encoding/json"
	"fmt"
)

// ? Todo関係
// !---------------------------------------------------------------------------------
// response
type Todos struct {
	Todo []TodoResponse `json:"todos"`
}
type TodoResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	IsDone    bool   `json:"isDone"`
}

type TodoMethod interface {
	CheckTitle() error
}

// request
type TodoCreateRequest struct {
	Title string `json:"title"`
}

func (t TodoCreateRequest) CheckTitle() error {
	return validateString(t.Title, "TodoTitle")
}

type TodoUpdateRequest struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	IsDone    bool   `json:"isDone"`
}

func (t TodoUpdateRequest) CheckTitle() error {
	return validateString(t.Title, "TodoTitle")
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
	IsDone    bool   `gorm:"default:false"`
}

// !---------------------------------------------------------------------------------
// ? validate
// !---------------------------------------------------------------------------------

// (Title or UserName, Type)
func validateString(s, t string) error {
	if s == "" || len(s) > 255 {
		return fmt.Errorf("%sがないか、255文字以上です", t)
	}
	return nil
}

// !---------------------------------------------------------------------------------
// ? error
// !---------------------------------------------------------------------------------
// error response

type ErrResponse struct {
	Error []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Tittle string `json:"tittle"`
	Detail string `json:"detail"`
}

//!---------------------------------------------------------------------------------

//? userregister
//!---------------------------------------------------------------------------------
//request

//response

//usertable db migrate

//error response

//error message
//!---------------------------------------------------------------------------------

//? login JWT
//!---------------------------------------------------------------------------------

//request

//response

//error response

//!---------------------------------------------------------------------------------

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
