package todo_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-webapi-todo/controller/todo"
	"github.com/gs223gs/go-webapi-todo/structs"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	todo.V1RestTodo(r, db)
	return r
}
func testTodoRequest(t *testing.T, r *gin.Engine, method, path, todoJSON string, expectedStatus int, expectedBody string) {
	req, _ := http.NewRequest(method, path, strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, expectedStatus, w.Code)
	if expectedBody != "" {
		assert.Contains(t, w.Body.String(), expectedBody)
	}
}
func TestGetTodos(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{})
	db.Create(&structs.Todos{Id: 1, Title: "Test Todo", Content: "Test Content", Category_Id: 1, Is_Done: false})

	r := setupRouter(db)
	testTodoRequest(t, r, "GET", "/v1/rest/todo", "", http.StatusOK, "Test Todo")
}
func TestPostTodo(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{}, &structs.Categories{})

	// カテゴリを追加
	db.Create(&structs.Categories{Id: 1, Category: "Test Category"})

	r := setupRouter(db)

	// テストケースの定義
	tests := []struct {
		name         string
		todoJSON     string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "正常なリクエスト",
			todoJSON:     `{"Title": "New Todo", "Content": "New Content", "Category_Id": 1}`,
			expectedCode: http.StatusOK,
			expectedBody: "追加完了",
		},
		{
			name:         "Titleがない場合",
			todoJSON:     `{"Content": "New Content", "Category_Id": 1}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
		// 他のテストケースも同様に追加
	}

	// テストケースの実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTodoRequest(t, r, "POST", "/v1/rest/todo", tt.todoJSON, tt.expectedCode, tt.expectedBody)
		})
	}
}

func TestPutTodo(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{}, &structs.Categories{})

	// 初期データ設定
	db.Create(&structs.Categories{Id: 1, Category: "Test Category"})
	db.Create(&structs.Todos{Id: 1, Title: "Old Todo", Content: "Old Content", Category_Id: 1, Is_Done: false})

	r := setupRouter(db)

	tests := []struct {
		name         string
		todoJSON     string
		expectedCode int
		expectedBody string
		checkDB      bool // DBの検証が必要な場合true
	}{
		{
			name:         "更新完了",
			todoJSON:     `{"Id":1,"Title": "Updated Todo", "Content": "Updated Content", "Category_Id": 1, "Is_Done": true}`,
			expectedCode: http.StatusOK,
			expectedBody: "更新完了",
			checkDB:      true,
		},
		{
			name:         "存在しないCategory_Id",
			todoJSON:     `{"Id":1,"Title": "Updated Todo", "Content": "Updated Content", "Category_Id": 999, "Is_Done": true}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "カテゴリが存在しません",
		},
		{
			name:         "Titleがない場合",
			todoJSON:     `{"Id":1,"Content": "Updated Content", "Category_Id": 1, "Is_Done": true}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Todo名がありません",
		},
		{
			name:         "Contentがnullの場合",
			todoJSON:     `{"Id":1,"Title": "Updated Todo", "Category_Id": 1, "Is_Done": true}`,
			expectedCode: http.StatusOK,
			expectedBody: "更新完了",
		},
		{
			name:         "Is_Doneがない場合",
			todoJSON:     `{"Id":1,"Title": "Updated Todo", "Content": "Updated Content", "Category_Id": 1}`,
			expectedCode: http.StatusOK,
			expectedBody: "更新完了",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTodoRequest(t, r, "PUT", "/v1/rest/todo", tt.todoJSON, tt.expectedCode, tt.expectedBody)

			if tt.checkDB {
				var updatedTodo structs.Todos
				if err := db.First(&updatedTodo, 1).Error; err != nil {
					t.Fatalf("更新されたTodoが見つかりません: %v", err)
				}
				assert.Equal(t, "Updated Todo", updatedTodo.Title)
				assert.Equal(t, "Updated Content", updatedTodo.Content)
				assert.Equal(t, uint(1), updatedTodo.Category_Id)
				assert.Equal(t, true, updatedTodo.Is_Done)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{})
	db.Create(&structs.Todos{Id: 1, Title: "Delete Todo", Content: "Delete Content", Category_Id: 1, Is_Done: false})

	r := setupRouter(db)

	tests := []struct {
		name         string
		todoJSON     string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "正常な削除",
			todoJSON:     `{"Id":1}`,
			expectedCode: http.StatusOK,
			expectedBody: "消去完了",
		},
		{
			name:         "存在しないTodoの削除",
			todoJSON:     `{"Id":999}`,
			expectedCode: http.StatusNotFound,
			expectedBody: "Todoが存在しません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTodoRequest(t, r, "DELETE", "/v1/rest/todo", tt.todoJSON, tt.expectedCode, tt.expectedBody)
		})
	}
}
