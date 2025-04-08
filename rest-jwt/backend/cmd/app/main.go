package main

import (
	"encoding/json"
	"fmt"
)

// ? Todo関係
// !---------------------------------------------------------------------------------
// response
type Todos struct {
	Todo []Todo `json:"todos"`
}

// JSON responseとDBのstruct
type Todo struct {
	Id     int    `json:"id" gorm:"primary_key;auto_increment"`
	Title  string `json:"title" gorm:"size:255"`
	IsDone bool   `json:"isDone" gorm:"default:false"`
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

func (t Todo) CheckTitle() error {
	return validateString(t.Title, "TodoTitle")
}

func (t Todo) CheckId() error {
	if t.Id <= 0 {
		return fmt.Errorf("Todo IDがないか、0以下です")
	}
	return nil
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

// ? userregister
// !---------------------------------------------------------------------------------
// request
type User struct {
	Id       int    `json:"id" gorm:"primary_key;auto_increment"`
	UserName string `json:"userName" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
}

//response

//!---------------------------------------------------------------------------------

//? JWT
//!---------------------------------------------------------------------------------
//request
//header payload signature
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub string `json:"sub"`
	Exp int    `json:"exp"`
}

type LoginRequest struct {
	Header Header `json:"header"`
	Payload Payload `json:"payload"`
	Signature string `json:"signature"`
}










//response

//!---------------------------------------------------------------------------------

//?JWT
//!---------------------------------------------------------------------------------
//request

func main() {
	var CreateRequest TodoMethod = &TodoCreateRequest{""}
	// var UpdateRequest TodoMethod = &TodoUpdateRequest{1, "a", false}
	UpdateRequest := Todo{0, "", false}

	if err := CreateRequest.CheckTitle(); err != nil {
		fmt.Println(err)
	}

	if err := UpdateRequest.CheckTitle(); err != nil {
		fmt.Println(err)
	}

	if err := UpdateRequest.CheckId(); err != nil {
		fmt.Println(err)
	}

	response := Todo{1, "test", false}
	todolist := Todos{Todo: []Todo{response, response}}

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
