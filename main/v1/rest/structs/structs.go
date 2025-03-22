package structs

import "time"

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
