package controller
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/gs223gs/go-webapi-todo/structs"
)



func V1RestCategory(r *gin.Engine, db *gorm.DB){
	r.GET("/v1/rest/category", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
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
		c.JSON(200, gin.H{
			"message": "hi PUT postman!",
		})
	})

	r.DELETE("/v1/rest/category", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi DELETE postman!",
		})
	})
}