package todo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-api-todo/controller/validation"
	"github.com/gs223gs/go-api-todo/structs"
	"gorm.io/gorm"
)

func V1RestTodo(r *gin.Engine, db *gorm.DB) {
	r.GET("/v1/rest/todo", func(c *gin.Context) {
		var todos []structs.Todos
		db.Select("Id", "Title", "Content", "Category_Id", "Due", "Is_Done", "Created_at", "Updated_at").Find(&todos)

		var todosResponse []structs.TodosResponse
		for _, todo := range todos {

			todosResponse = append(todosResponse, structs.TodosResponse{
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
				Updated_at: todo.Updated_at.Format(time.RFC3339),
			})
		}

		c.JSON(http.StatusOK, todosResponse)
	})

	r.POST("/v1/rest/todo", func(c *gin.Context) {
		var todo structs.Todos

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validate = map[string]any{"TodoTitle": todo.Title, "CategoryID": todo.Category_Id}
		if err := validation.Check(validate, db); len(err) != 0 {
			c.JSON(http.StatusBadRequest, gin.H(validation.Conv(err)))
			return
		}

		if err := db.Create(&todo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "登録に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"messege": "追加完了"})
	})

	////REST原則に基づいて PUTを実装し直す
	r.PUT("/v1/rest/todo", func(c *gin.Context) {
		var updateTodo structs.Todos
		if err := c.ShouldBindJSON(&updateTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validate = map[string]any{"TodoID": updateTodo.Id, "TodoTitle": updateTodo.Title, "CategoryID": updateTodo.Category_Id}
		if err := validation.Check(validate, db); len(err) != 0 {
			c.JSON(http.StatusBadRequest, gin.H(validation.Conv(err)))
			return
		}

		fmt.Println(updateTodo)

		/*
			既存Todoの取得
			これにしないとupdated_at等が更新できない
			理由:PUTは明示的に全てのカラムを送らなければいけない？
			部分的な更新はPATCHで行う？
		*/
		var existingTodo structs.Todos

		existingTodo.Title = updateTodo.Title
		existingTodo.Content = updateTodo.Content
		existingTodo.Category_Id = updateTodo.Category_Id
		existingTodo.Is_Done = updateTodo.Is_Done
		existingTodo.Due = updateTodo.Due
		existingTodo.Updated_at = time.Now()

		if err := db.Save(&existingTodo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "更新完了"})
	})

	r.DELETE("/v1/rest/todo", func(c *gin.Context) {
		var todo structs.Todos

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := todo.CheckID(db); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Delete(&todo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "削除に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messege": "消去完了"})
	})
}
