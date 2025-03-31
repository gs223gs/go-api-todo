package structs

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todos struct {
	Id          uint       `gorm:"primary_key;autoIncrement"`
	Title       string     `gorm:"size:255"`
	Content     string     `gorm:"text"`
	Category_Id uint       `gorm:"size:100"`
	Is_Done     bool       `gorm:"default:false"`
	Due         *time.Time `gorm:"default:NULL"`
	Created_at  time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at  time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	Category    Categories `gorm:"foreignKey:Category_Id;references:Id;constraint:OnDelete:CASCADE"`
}

/*
	validation checkロジック
	request => BindJSON => todo.メソッド() の順でこのメソッドに辿り着く
	BindJSONで型チェックが行われるため，型自体のチェックは必要がない.
	IDの存在確認やstrngのlengthチェックを行う
*/

func (t Todos) CheckID(db *gorm.DB) error {

	if err := db.First(&t, t.Id).Error; err != nil {
		return fmt.Errorf("Todoがありません")
	}
	return nil
}

func (t Todos) CheckTitle() error {
	maxTitleLength := 255
	if t.Title == "" {
		return fmt.Errorf("Titleがありません")
	}
	if len(t.Title) > maxTitleLength {
		return fmt.Errorf("Titleが長すぎます")
	}
	return nil
}

func (t Todos) CheckCategoryId(db *gorm.DB) error {
	var categories Categories
	if err := db.First(&categories, t.Category_Id).Error; err != nil {
		return fmt.Errorf("カテゴリが存在しません")
	}
	return nil
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
	Updated_at  string
}

/*
todoリスト
バリデーションチェックのメソッド作成
when: 2025-03-28
where:
who: T.Miura
what: todos & categories struct のバリデーションチェック
why: validation packageが美しくない
how:

test項目 {
	.todosメソッドについて {
		id T = idに紐づくtodoが存在する nil F = "todoが存在しません"
		title T = titleが""（空文字）ではない nil F = "Titleがありません"
		Category_id T = Categories table に category_idに紐づくカテゴリが存在する nil F = "カテゴリが存在しません"
		Is_Done T = bool型以外が来ていない nil F = "不正な入力値です"
		Due T =  yyyy-mm-dd の形式になっているかどうか nil F =  "規格があっていません"
	}
	categoriesメソッドについて{
		id T = idに紐づくcategoryが存在するか nil F = "カテゴリが存在しません"
		category = categoryが""(空文字)ではないか nil F = "カテゴリ名がありません"
	}
}



methodについて
CheckID
CheckTitle
CheckCategoryId
CheckIsDone 必要なし
CheckDue　必要なし


使い方
var todo Todos
db gorm
if err ....(JSON バインド)

var errorMessege = map[string]string

[idをチェックする場合]
if err := todo.CheckId(db); err != nil {
	errorMessege[todoID] = err.Error()
}

[error の時の response]
if errorMessege != nil {
gin.H(errorMessege)
}



設計について
エラー時にJSONでレスポンスする

{
	"errors": [
		{
			"field": "title",
			"message": "titleがありません"
		},
		{
			"field": "category_id",
			"message": "カテゴリが存在しません"
		}
	]
}
*/
