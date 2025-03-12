package structs

import "time"

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
