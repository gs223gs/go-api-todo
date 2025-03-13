package category

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
		var Category structs.Categories
		if err := c.ShouldBindJSON(&Category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&Category)
		c.JSON(http.StatusOK, Category)
	})

	r.PUT("/v1/rest/category", func(c *gin.Context) {
		var categories structs.Categories

		if err := c.ShouldBindJSON(&categories); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.First(&categories, categories.Id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "カテゴリが存在しません"})
			return
		}
		fmt.Printf("%T %v----------------------------------------", categories, categories)
		db.Save(&categories)
		c.JSON(http.StatusOK, categories)
	})

	r.DELETE("/v1/rest/category", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi DELETE postman!",
		})
	})
}
