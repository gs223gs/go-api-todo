package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-api-todo/controller/category"
	"github.com/gs223gs/go-api-todo/controller/todo"
	"github.com/gs223gs/go-api-todo/db"
)

func main() {

	Database := db.ConnctionDb()
	db.InitDB(Database)
	r := gin.Default()

	todo.V1RestTodo(r, Database)
	category.V1RestCategory(r, Database)

	r.Run(":8080")
}
