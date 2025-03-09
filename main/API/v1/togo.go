package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"models/v1/models"
)

func v1RestTodo(r *gin.Engine, db *gorm.DB) {
	r.GET("/v1/rest/todo", func(c *gin.Context) {
		var todos []models.Todos
		db.Select("Id", "Title", "Content", "Category_Id", "Due", "Is_Done", "Created_at").Find(&todos)

		var todosResponse []models.TodosResponse
		for _, todo := range todos {

			todosResponse = append(todosResponse, models.TodosResponse{
				Id:          todo.Id,
				Title:       todo.Title,
				Content:     todo.Content,
				Category_id: todo.Category_Id,
				Is_Done:     todo.Is_Done,
				Due: func() string {
					if todo.Due != nil {
						return todo.Due.Format(time.RFC3339)
					}
					return ""
				}(),
				Created_at: todo.Created_at.Format(time.RFC3339),
			})
		}

		c.JSON(http.StatusOK, todosResponse)
	})

	r.POST("/v1/rest/todo", func(c *gin.Context) {
		var todo models.odos
		var categories models.Categories
		// JSONからデータをバインド
		if err := c.ShouldBindJSON(&todo); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// JSONからCategory_Idを取得
		category := todo.Category_Id

		// Categoryの存在を確認
		if err := db.First(&categories, category).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "カテゴリが存在しません"})
			return
		}

		db.Create(&todo)
		c.JSON(http.StatusOK, gin.H{"messege": "追加完了"})
	})

	r.PUT("/v1/rest/todo", func(c *gin.Context) {
		var todo models.Todos
		var categories models.Categories
		id := c.Param("Id")
		category := c.Param("Category_id")

		if err := db.First(&categories, category).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "カテゴリが存在しません"})
			return
		}

		if err := db.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todoが存在しません"})
			return
		}

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&todo)
		c.JSON(http.StatusOK, todo)
	})

	r.DELETE("/v1/rest/todo", func(c *gin.Context) {
		var todo models.Todos

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.First(&todo, todo.Id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "todoが存在しません"})
			return
		}

		db.Delete(&todo)
		c.JSON(http.StatusOK, gin.H{"messege": "消去完了"})
	})
}
