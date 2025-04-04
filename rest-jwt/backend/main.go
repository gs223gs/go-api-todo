package main

import (
	"fmt"
)
//response
type Todos struct {
	Todo []TodoResponse `json:"todos"`
}
type TodoResponse struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}


type ErrResponse struct {
	Message map[string]any `json:"message"`
}

type TodoMethod interface {
	CheckTitle() error
}



//request
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
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

//gorm db mygrate
type TodoTable struct {
	ID int `gorm:"primary_key;auto_increment"`
	Title string `gorm:"size:255"`
	Completed bool `gorm:"default:false"`
}

func main() {
	var CreateRequest TodoMethod = &TodoCreateRequest{"test"}
	if err := CreateRequest.CheckTitle(); err != nil {
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