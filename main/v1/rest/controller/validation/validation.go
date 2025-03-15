package validation

import (
	"fmt"
	"strconv"

	"github.com/gs223gs/go-webapi-todo/structs"
	"gorm.io/gorm"
)

// ! package名がvalidationのため，関数名をわざと短くしている

// .TODOIDがDBに存在するかチェック
func TodoID(id int, db *gorm.DB) error {
	var todo structs.Todos
	if err := db.First(&todo, int(id)).Error; err != nil {
		return fmt.Errorf("Todoがありません")
	}
	return nil
}

func CategoryID(id int, db *gorm.DB) error {
	var categories structs.Categories
	if err := db.First(&categories, id).Error; err != nil {
		return fmt.Errorf("カテゴリがありません")
	}
	return nil
}

func ContentType(contentType, supportType string) error {
	if contentType != supportType {
		return fmt.Errorf("サポートされていないメディアタイプです．")
	}
	return nil
}

func TodoTitle(Title string) error {
	if Title == "" {
		return fmt.Errorf("Todo名がありません")
	}
	return nil
}

func Check(s map[string]string, db *gorm.DB) (result map[string]error) {
	result = make(map[string]error)
	for k, v := range s {
		switch k {
		case "TodoID":
			if id, err := strconv.Atoi(v); err == nil {
				fmt.Println(id)
				if err := TodoID(id, db); err != nil {
					result["TodoID"] = err
				}
			} else {
				result["TodoID"] = fmt.Errorf("無効なTodoIDです")
			}

		case "TodoTitle":
			if err := TodoTitle(v); err != nil {
				result["TodoTitle"] = err
			}
		case "CategoryID":
		case "Content-Type":
		}
	}
	return result
}
