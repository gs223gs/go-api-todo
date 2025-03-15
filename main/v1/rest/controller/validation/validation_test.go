package validation_test

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gs223gs/go-webapi-todo/controller/validation"
	"github.com/gs223gs/go-webapi-todo/structs"
	_ "github.com/mattn/go-sqlite3"
)

func TestCheck(t *testing.T) {
	var db *gorm.DB = nil

	tests := []struct {
		name  string
		input map[string]string
		// 期待するエラーは、キーとエラーメッセージのマップとして定義
		want map[string]string
	}{
		{
			name: "TodoTitleが存在する場合（正常）",
			input: map[string]string{
				"TodoTitle": "有効なタイトル",
			},
			want: map[string]string{}, // エラーなし
		},
		{
			name: "TodoTitleが空の場合（エラー発生）",
			input: map[string]string{
				"TodoTitle": "",
			},
			want: map[string]string{
				"TodoTitle": "Todo名がありません",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validation.Check(tt.input, db)
			if len(got) != len(tt.want) {
				t.Errorf("期待するエラー数は %d でしたが、実際は %d でした", len(tt.want), len(got))
			}
			for key, expectedMsg := range tt.want {
				err, exists := got[key]
				if !exists {
					t.Errorf("キー %s のエラーが返却されていません", key)
					continue
				}
				if err.Error() != expectedMsg {
					t.Errorf("キー %s のエラー: 期待値 %q、実際の値 %q", key, expectedMsg, err.Error())
				}
			}
		})
	}
}

func TestTodoID(t *testing.T) {
	// テストデータベースのセットアップ
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("データベース接続に失敗しました: %v", err)
	}

	// テーブルの作成
	err = db.AutoMigrate(&structs.Todos{})
	if err != nil {
		t.Fatalf("マイグレーションに失敗しました: %v", err)
	}

	// テストデータの作成
	testTodo := structs.Todos{
		Id:          1,
		Title:       "Test Todo",
		Content:     "Test Content",
		Category_Id: 1,
	}
	db.Create(&testTodo)

	tests := []struct {
		name    string
		id      int
		wantErr bool
		errMsg  string
	}{
		{
			name:    "存在するID",
			id:      1,
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "存在しないID",
			id:      999,
			wantErr: true,
			errMsg:  "Todoがありません",
		},
		{
			name:    "不正な入力（負の値）",
			id:      -1,
			wantErr: true,
			errMsg:  "Todoがありません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.TodoID(tt.id, db)

			// エラーの有無をチェック
			if (err != nil) != tt.wantErr {
				t.Errorf("TodoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// エラーメッセージをチェック
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("TodoID() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}

func TestCategoryID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("データベース接続に失敗しました: %v", err)
	}

	// テーブルの作成
	err = db.AutoMigrate(&structs.Categories{})
	if err != nil {
		t.Fatalf("マイグレーションに失敗しました: %v", err)
	}

	// テストデータの作成
	testCategory := structs.Categories{
		Id:       1,
		Category: "テストカテゴリ",
	}
	db.Create(&testCategory)

	tests := []struct {
		name    string
		id      int
		wantErr bool
		errMsg  string
	}{
		{
			name:    "存在するID",
			id:      1,
			wantErr: false,
			errMsg:  "",
		},
		{
			name:    "存在しないID",
			id:      999,
			wantErr: true,
			errMsg:  "カテゴリがありません",
		},
		{
			name:    "不正な入力（負の値）",
			id:      -1,
			wantErr: true,
			errMsg:  "カテゴリがありません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.CategoryID(tt.id, db)

			// エラーの有無をチェック
			if (err != nil) != tt.wantErr {
				t.Errorf(" CategoryID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// エラーメッセージをチェック
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf(" CategoryID() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
