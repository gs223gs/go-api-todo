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
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("データベースの作成に失敗しました: %v", err)
	}

	// テーブルの作成
	if err := db.AutoMigrate(&structs.Todos{}, &structs.Categories{}); err != nil {
		t.Fatalf("テーブルの作成に失敗しました: %v", err)
	}
	// テストデータの作成
	categories := []structs.Categories{
		{Id: 1, Category: "カテゴリ1"},
		{Id: 2, Category: "カテゴリ2"},
		{Id: 3, Category: "カテゴリ3"},
	}

	todos := []structs.Todos{
		{Id: 1, Title: "Todo1", Content: "内容1", Category_Id: 1},
		{Id: 2, Title: "Todo2", Content: "内容2", Category_Id: 2},
		{Id: 3, Title: "Todo3", Content: "内容3", Category_Id: 3},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			t.Fatalf("カテゴリの作成に失敗しました: %v", err)
		}
	}

	for _, todo := range todos {
		if err := db.Create(&todo).Error; err != nil {
			t.Fatalf("Todoの作成に失敗しました: %v", err)
		}
	}

	tests := []struct {
		name  string
		input map[string]string
		want  map[string]string
	}{
		{
			name: "TodoTitleが存在する場合（正常）",
			input: map[string]string{
				"TodoTitle": "有効なタイトル",
				"TodoID":    "1",
			},
			want: map[string]string{},
		},
		{
			name: "無効なTodoIDが存在する場合",
			input: map[string]string{
				"TodoID": "無効なID",
			},
			want: map[string]string{
				"TodoID": "無効なTodoIDです",
			},
		},
		{
			name: "存在しないTodoIDが存在する場合",
			input: map[string]string{
				"TodoID": "999",
			},
			want: map[string]string{
				"TodoID": "Todoがありません",
			},
		},
		{
			name: "TodoTitleが空の場合",
			input: map[string]string{
				"TodoTitle": "",
			},
			want: map[string]string{
				"TodoTitle": "Todo名がありません",
			},
		},
		{
			name: "無効なCategoryIDが存在する場合",
			input: map[string]string{
				"CategoryID": "無効なID",
			},
			want: map[string]string{
				"CategoryID": "無効なCategoryIDです",
			},
		},
		{
			name: "存在しないCategoryIDが存在する場合",
			input: map[string]string{
				"CategoryID": "999",
			},
			want: map[string]string{
				"CategoryID": "カテゴリがありません",
			},
		},
		{
			name: "サポートされていないContent-Typeが存在する場合",
			input: map[string]string{
				"Content-Type": "unsupported/type",
				"supportType":  "supported/type",
			},
			want: map[string]string{
				"Content-Type": "サポートされていないメディアタイプです．",
			},
		},
		{
			name: "supportTypeが存在しない場合",
			input: map[string]string{
				"Content-Type": "supported/type",
			},
			want: map[string]string{
				"supportType": "内部エラー",
			},
		},
	}

	for _, tt := range tests {
		tt := tt // ループ変数のスコープを固定
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // テストを並行して実行
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

func TestContentType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		supportType string
		wantErr     bool
		errMsg      string
	}{
		{
			name:        "サポートされているタイプ",
			contentType: "application/json",
			supportType: "application/json",
			wantErr:     false,
			errMsg:      "",
		},
		{
			name:        "サポートされていないタイプ",
			contentType: "application/xml",
			supportType: "application/json",
			wantErr:     true,
			errMsg:      "サポートされていないメディアタイプです．",
		},
		{
			name:        "空のContent-Type",
			contentType: "",
			supportType: "application/json",
			wantErr:     true,
			errMsg:      "サポートされていないメディアタイプです．",
		},
		{
			name:        "大文字小文字の違い",
			contentType: "APPLICATION/JSON",
			supportType: "application/json",
			wantErr:     true,
			errMsg:      "サポートされていないメディアタイプです．",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ContentType(tt.contentType, tt.supportType)

			// エラーの有無をチェック
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// エラーメッセージをチェック
			if tt.wantErr && err.Error() != tt.errMsg {
				t.Errorf("ContentType() error message = %v, want %v", err.Error(), tt.errMsg)
			}
		})
	}
}
