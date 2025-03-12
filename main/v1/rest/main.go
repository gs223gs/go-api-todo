package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-webapi-todo/controller"
	"github.com/gs223gs/go-webapi-todo/db"
)

func main() {

	Database := db.ConnctionDb()
	db.InitDB(Database)
	r := gin.Default()

	controller.V1RestTodo(r, Database)
	controller.V1RestCategory(r, Database)

	r.Run(":8080")
}
