package todo

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-webapi-todo/controller/validation"
	"github.com/gs223gs/go-webapi-todo/structs"
	"gorm.io/gorm"
)

func V1RestTodo(r *gin.Engine, db *gorm.DB) {
	r.GET("/v1/rest/todo", func(c *gin.Context) {
		var todos []structs.Todos
		db.Select("Id", "Title", "Content", "Category_Id", "Due", "Is_Done", "Created_at").Find(&todos)

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

	//REST原則に基づいて PUTを実装し直す
	r.PUT("/v1/rest/todo", func(c *gin.Context) {
		var todo structs.Todos
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validate = map[string]any{"TodoID": todo.Id, "TodoTitle": todo.Title, "CategoryID": todo.Category_Id}
		if err := validation.Check(validate, db); len(err) != 0 {
			c.JSON(http.StatusBadRequest, gin.H(validation.Conv(err)))
			return
		}

		var existingTodo structs.Todos
		if err := db.First(&existingTodo, todo.Id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todoが見つかりません"})
			return
		}

		timenow := time.Now()

		existingTodo.Title = todo.Title
		existingTodo.Content = todo.Content
		existingTodo.Category_Id = todo.Category_Id
		existingTodo.Is_Done = todo.Is_Done
		existingTodo.Due = todo.Due
		existingTodo.Updated_at = timenow

		if err := db.Save(&existingTodo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messege": "更新完了"})
	})

	r.DELETE("/v1/rest/todo", func(c *gin.Context) {
		var todo structs.Todos

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validate = map[string]any{"TodoID": todo.Id}
		if err := validation.Check(validate, db); len(err) != 0 {
			c.JSON(http.StatusBadRequest, gin.H(validation.Conv(err)))
			return
		}

		if err := db.Delete(&todo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "削除に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"messege": "消去完了"})
	})
}
