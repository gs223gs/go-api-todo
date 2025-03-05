package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createDSN(host, user, password, dbname, port string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbname, port)
}

type Todos struct {
	Id          uint      `gorm:"primary_key"`
	Title       string    `gorm:"size:255"`
	Content     string    `gorm:"text"`
	Category_Id int       `gorm:"size:100;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Category_Id;references:Id"`
	Is_Done     bool      `gorm:"default:false"`
	Created_at  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Update_at   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type Categories struct {
	Id       uint   `gorm:"primary_key"`
	Category string `gorm:"site:255"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env ファイルがありません")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")

	dsn := createDSN(dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DBに接続できません", err)
	}

	db.AutoMigrate(&Categories{}, &Todos{})
	r := gin.Default()

	//###############################################################################
	r.GET("/v1/rest/todo", func(c *gin.Context) {
		var todos []Todos
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	r.POST("/v1/rest/todo", func(c *gin.Context) {
		var todo Todos
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&todo)
		c.JSON(http.StatusOK, todo)
	})

	r.PUT("/v1/rest/todo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi PUT postman!",
		})
	})

	r.DELETE("/v1/rest/todo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi DELETE postman!",
		})
	})
	//###############################################################################

	//###############################################################################
	r.GET("/v1/rest/category", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/v1/rest/category", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi POST postman!",
		})
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

	//###############################################################################
	r.Run(":8080")
}
