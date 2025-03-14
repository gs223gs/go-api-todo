package validation

import (
	"fmt"

	"gorm.io/gorm"
)

// ! package名がvalidationのため，関数名をわざと短くしている
func TodoID() {

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

/*
check(s スライス ) (map[string]stringerr) {
result map[string]string
swich
	TodoID
		if check {
		,todoid :="TodoID"
			err := checkfunc(TodoID); err != nil{
			result[todoid] = err
			}
		}
	CategoryID
		if check{
			return err
		}
	ContentType
		if check{
			return err
		}
	return result
}
*/
