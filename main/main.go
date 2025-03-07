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
	Id          uint       `gorm:"primary_key;autoIncrement"`
	Title       string     `gorm:"size:255"`
	Content     string     `gorm:"text"`
	Category_Id uint       `gorm:"size:100;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:Category_Id;references:Id"` //CategoriesのidでforeingKey
	Is_Done     bool       `gorm:"default:false"`
	Due         *time.Time `gorm:"default:NULL"` //ポインタ使用でNULL許容
	Created_at  time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Update_at   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Category    Categories `gorm:"foreignKey:Category_Id;references:Id"`
}

type Categories struct {
	Id       uint   `gorm:"primary_key;autoIncrement"`
	Category string `gorm:"size:255"`
}

type TodosResponse struct {
	Id          uint
	Title       string
	Content     string
	Category_id uint
	Is_Done     bool
	Due         string
	Created_at  string
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

		fmt.Println(todosResponse)

		c.JSON(http.StatusOK, todosResponse)
	})

	r.POST("/v1/rest/todo", func(c *gin.Context) {
		var todo Todos
		var categories Categories
		fmt.Println("a:", todo)
		fmt.Println("C :", c)
		// JSONからデータをバインド
		if err := c.ShouldBindJSON(&todo); err != nil {
			fmt.Println("Error: ", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// JSONからCategory_Idを取得
		category := todo.Category_Id
		fmt.Println("Category ID:", category)

		// Categoryの存在を確認
		if err := db.First(&categories, category).Error; err != nil {
			fmt.Println("Error: category not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
			return
		}

		// TodoをDBに作成
		db.Create(&todo)
		c.JSON(http.StatusOK, gin.H{"messege": "Create Success"})
	})

	r.PUT("/v1/rest/todo", func(c *gin.Context) {
		var todo Todos
		var categories Categories
		id := c.Param("Id")
		category := c.Param("Category_id")

		fmt.Println("ID:", id)
		fmt.Println("Category ID:", category)

		if err := db.First(&categories, category).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}

		if err := db.First(&todo, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
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

	//###############################################################################
	r.Run(":8080")
}
