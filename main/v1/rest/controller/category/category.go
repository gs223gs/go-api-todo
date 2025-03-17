package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-webapi-todo/controller/validation"
	"github.com/gs223gs/go-webapi-todo/structs"
	"gorm.io/gorm"
)

func V1RestCategory(r *gin.Engine, db *gorm.DB) {
	r.GET("/v1/rest/category", func(c *gin.Context) {

		var Categories []structs.Categories
		db.Select("Id", "Category").Find(&Categories)
		c.JSON(http.StatusOK, Categories)
	})

	r.POST("/v1/rest/category", func(c *gin.Context) {
		var categories structs.Categories

		if err := c.ShouldBindJSON(&categories); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validate = map[string]any{"CategoryTitle": categories.Category}
		if err := validation.Check(validate, db); len(err) != 0 {
			c.JSON(http.StatusBadRequest, gin.H(validation.Conv(err)))
			return
		}

		
		if err := db.Create(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "登録に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, categories)
	})

	r.PUT("/v1/rest/category", func(c *gin.Context) {
		var categories structs.Categories

		if err := c.ShouldBindJSON(&categories); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var validate = map[string]any{"CategoryID": categories.Id, "CategoryTitle": categories.Category}
		if err := validation.Check(validate, db); len(err) != 0 {
			c.JSON(http.StatusBadRequest, gin.H(validation.Conv(err)))
			return
		}

		if err := db.Save(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "登録に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, categories)
	})

	r.DELETE("/v1/rest/category", func(c *gin.Context) {
		var categories structs.Categories

		var validate = map[string]any{"CategoryID": categories.Id}
		if err := validation.Check(validate, db); len(err) != 0 {
			c.JSON(http.StatusBadRequest, gin.H(validation.Conv(err)))
			return
		}

		if err := db.Delete(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "削除に失敗しました"})
			return
		}
	})
}
