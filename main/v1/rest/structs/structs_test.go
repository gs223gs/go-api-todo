package structs_test

import (
	"testing"

	"github.com/gs223gs/go-api-todo/structs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// SQLiteを使用してテスト用のインメモリデータベースを作成
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// テーブルを作成
	db.AutoMigrate(&structs.Todos{}, &structs.Categories{})

	// テスト用のデータを挿入
	db.Create(&structs.Categories{Id: 1, Category: "仕事"})
	db.Create(&structs.Todos{Id: 1, Title: "タスク1", Category_Id: 1, Is_Done: false})

	return db
}

func TestCheckID(t *testing.T) {
	db := setupTestDB()

	tests := []struct {
		name    string
		todo    structs.Todos
		wantErr bool
		errMsg  string
	}{
		{
			name:    "存在するID",
			todo:    structs.Todos{Id: 1},
			wantErr: false,
		},
		{
			name:    "存在しないID",
			todo:    structs.Todos{Id: 999},
			wantErr: true,
			errMsg:  "Todoがありません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.todo.CheckID(db)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckID() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("CheckID() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestCheckTitle(t *testing.T) {
	tests := []struct {
		name    string
		todo    structs.Todos
		wantErr bool
		errMsg  string
	}{
		{
			name:    "有効なタイトル",
			todo:    structs.Todos{Title: "有効なタイトル"},
			wantErr: false,
		},
		{
			name:    "空のタイトル",
			todo:    structs.Todos{Title: ""},
			wantErr: true,
			errMsg:  "Titleがありません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.todo.CheckTitle()
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckTitle() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("CheckTitle() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestCheckCategoryId(t *testing.T) {
	db := setupTestDB()

	tests := []struct {
		name    string
		todo    structs.Todos
		wantErr bool
		errMsg  string
	}{
		{
			name:    "存在するCategoryID",
			todo:    structs.Todos{Category_Id: 1},
			wantErr: false,
		},
		{
			name:    "存在しないCategoryID",
			todo:    structs.Todos{Category_Id: 999},
			wantErr: true,
			errMsg:  "カテゴリが存在しません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.todo.CheckCategoryId(db)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckCategoryId() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("CheckCategoryId() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
