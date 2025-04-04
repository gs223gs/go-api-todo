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

//request
type TodoCreateRequest struct {
	Title string `json:"title"`
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