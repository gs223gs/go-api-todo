package validation

import (
	"fmt"

	"github.com/gs223gs/go-api-todo/structs"
	"gorm.io/gorm"
)

// ! package名がvalidationのため，関数名をわざと短くしている

// .TODOIDがDBに存在するかチェック
func TodoID(id uint, db *gorm.DB) error {
	var todo structs.Todos
	if err := db.First(&todo, int(id)).Error; err != nil {
		return fmt.Errorf("Todoがありません")
	}
	return nil
}

func CategoryID(id uint, db *gorm.DB) error {
	var categories structs.Categories
	if err := db.First(&categories, id).Error; err != nil {
		return fmt.Errorf("カテゴリがありません")
	}
	return nil
}



func TodoTitle(Title string) error {
	if Title == "" {
		return fmt.Errorf("Todo名がありません")
	}
	return nil
}

func CategoryTitle(Title string) error {
	if Title == "" {
		return fmt.Errorf("Category名がありません")
	}
	return nil
}

/*
Key values ​​that can be used =>
TodoID,
TodoTitle,
CategoryID
*/
func Check(m map[string]any, db *gorm.DB) (result map[string]error) {
	result = make(map[string]error)
	for k, v := range m {
		switch k {
		case "TodoID":
			if id, ok := v.(uint); ok {
				if err := TodoID(id, db); err != nil {
					result["TodoID"] = err
				}
			} else {
				result["TodoID"] = fmt.Errorf("無効なTodoIDです")
			}
		case "TodoTitle":
			if str, ok := v.(string); ok {
				if err := TodoTitle(str); err != nil {
					result[k] = err
				}
			} else {
				result[k] = fmt.Errorf("無効なTodo名です")

			}
		case "CategoryTitle":
			if str, ok := v.(string); ok {
				if err := CategoryTitle(str); err != nil {
					result[k] = err
				}
			} else {
				result[k] = fmt.Errorf("無効なCategory名です")

			}
		case "CategoryID":
			if id, ok := v.(uint); ok {
				if err := CategoryID(id, db); err != nil {
					result["CategoryID"] = err
				}
			} else {
				result["CategoryID"] = fmt.Errorf("無効なCategoryIDです")
			}
		}
	}
	return
}

func Conv(prev map[string]error) (result map[string]any) {
	result = make(map[string]any)
	for k, v := range prev {
		result[k] = v.Error()
	}
	return
}
