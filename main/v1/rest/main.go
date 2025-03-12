package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-webapi-todo/controller"
	"github.com/gs223gs/go-webapi-todo/db"
	"github.com/gs223gs/go-webapi-todo/structs"
)

func main() {

	db := db.ConnctionDb()

	db.AutoMigrate(&structs.Categories{}, &structs.Todos{})

	r := gin.Default()

	controller.V1RestTodo(r, db)
	controller.V1RestCategory(r, db)

	r.Run(":8080")
}
