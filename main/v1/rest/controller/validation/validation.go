package validation

import (
	"fmt"

	"github.com/gs223gs/go-webapi-todo/structs"
	"gorm.io/gorm"
)

// ! package名がvalidationのため，関数名をわざと短くしている

//.TODOIDがDBに存在するかチェック
// 
func TodoID(id int, db *gorm.DB) error {
	var todo structs.Todos
	if err := db.First(&todo, int(id)).Error; err != nil {
		return fmt.Errorf("Todoがありません")
	}
	return nil

}

func CategoryID(db *gorm.DB, id int) {

}

func ContentType() {

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
		case "TodoId":
		case "TodoTitle":
			if err := TodoTitle(v); err != nil {
				result["TodoTitle"] = err
			}
		case "CategoryId":
		case "Content-Type":
		default:
			// デフォルトの処理をここに追加
		}
	}
	return result
}
