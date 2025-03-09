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

type Todos struct {
	Id          uint       `gorm:"primary_key;autoIncrement"`
	Title       string     `gorm:"size:255"`
	Content     string     `gorm:"text"`
	Category_Id uint       `gorm:"size:100;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Category_Id;references:Id"` //CategoriesのidでforeingKey
	Is_Done     bool       `gorm:"default:false"`
	Due         *time.Time `gorm:"default:NULL"` //ポインタ使用でNULL許容しています．
	Created_at  time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Update_at   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Category    Categories `gorm:"foreignKey:Category_Id;references:Id"`
}

type Categories struct {
	Id       uint   `gorm:"primary_key;autoIncrement"`
	Category string `gorm:"size:255"`
}

// get responseのため
type TodosResponse struct {
	Id          uint
	Title       string
	Content     string
	Category_id uint
	Is_Done     bool
	Due         string
	Created_at  string
}

func createDSN() (string, error) {
	err := godotenv.Load()

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", host, user, password, dbname, port)

	return dsn, err
}

func v1RestTodo(r *gin.Engine, db *gorm.DB) {
	r.GET("/v1/rest/todo", func(c *gin.Context) {
		var todos []Todos
		db.Select("Id", "Title", "Content", "Category_Id", "Due", "Is_Done", "Created_at").Find(&todos)

		var todosResponse []TodosResponse
		for _, todo := range todos {

			todosResponse = append(todosResponse, TodosResponse{
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
		var todo Todos
		var categories Categories
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
		var todo Todos
		var categories Categories
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
		var todo Todos

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

func v1RestCategory(r *gin.Engine, db *gorm.DB){
	r.GET("/v1/rest/category", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/v1/rest/category", func(c *gin.Context) {
		var Category Categories
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
func main() {

	dsn, err := createDSN()

	if err != nil {
		log.Fatal(".env ファイルがありません")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DBに接続できません", err)
	}

	db.AutoMigrate(&Categories{}, &Todos{})

	r := gin.Default()

	v1RestTodo(r, db)
	v1RestCategory(r,db)

	r.Run(":8080")
}
