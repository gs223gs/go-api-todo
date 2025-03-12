package db

import (
	"fmt"
	"log"
	"os"

	"github.com/gs223gs/go-webapi-todo/structs"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
func ConnctionDb() *gorm.DB {
	dsn, err := createDSN()

	if err != nil {
		log.Fatal(".env ファイルがありません")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DBに接続できません", err)
	}

	return db
}

func InitDB(db *gorm.DB) {
	db.Migrator().DropTable(&structs.Todos{}, &structs.Categories{})
	db.AutoMigrate(&structs.Categories{}, &structs.Todos{})
	var categories = []structs.Categories{
		{Category: "仕事"},
		{Category: "プライベート"},
		{Category: "勉強"},
	}

	for _, category := range categories {
		db.Create(&category)
	}

	var todos = []structs.Todos{
		{Title: "仕事のタスク1", Content: "仕事の内容1", Category_Id: 1, Is_Done: false},
		{Title: "プライベートのタスク1", Content: "プライベートの内容1", Category_Id: 2, Is_Done: false},
		{Title: "勉強のタスク1", Content: "勉強の内容1", Category_Id: 3, Is_Done: false},
	}

	for _, todo := range todos {
		db.Create(&todo)
	}
}
