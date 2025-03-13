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

func TestGetTodos(t *testing.T) {
	// SQLiteを使用したインメモリデータベースを作成
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{})

	// テストデータを挿入
	db.Create(&structs.Todos{Id: 1, Title: "Test Todo", Content: "Test Content", Category_Id: 1, Is_Done: false})

	r := setupRouter(db)

	req, _ := http.NewRequest("GET", "/v1/rest/todo", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Todo")
}

func TestPostTodo(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{}, &structs.Categories{})

	// カテゴリを追加
	db.Create(&structs.Categories{Id: 1, Category: "Test Category"})

	r := setupRouter(db)

	// 正常なリクエスト
	todoJSON := `{"Title": "New Todo", "Content": "New Content", "Category_Id": 1}`
	req, _ := http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "追加完了")

	// Titleがない場合
	todoJSON = `{"Content": "New Content", "Category_Id": 1}`
	req, _ = http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Category_Idがない場合
	todoJSON = `{"Title": "New Todo", "Content": "New Content"}`
	req, _ = http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Contentがnullの場合
	todoJSON = `{"Title": "New Todo", "Category_Id": 1}`
	req, _ = http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "追加完了")

	// Is_Doneが指定されていない場合
	todoJSON = `{"Title": "New Todo", "Content": "New Content", "Category_Id": 1}`
	req, _ = http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "追加完了")

	// 不正なJSONフォーマット
	todoJSON = `{"Title": "New Todo", "Content": "New Content", "Category_Id": 1`
	req, _ = http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// 不正なContent-Type
	todoJSON = `{"Title": "New Todo", "Content": "New Content", "Category_Id": 1}`
	req, _ = http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "text/plain")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnsupportedMediaType, w.Code)

	// 存在しないCategory_Idの場合
	todoJSON = `{"Title": "New Todo", "Content": "New Content", "Category_Id": 999}`
	req, _ = http.NewRequest("POST", "/v1/rest/todo", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPutTodo(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{}, &structs.Categories{})

	// カテゴリとTodoを追加
	db.Create(&structs.Categories{Id: 1, Category: "Test Category"})
	db.Create(&structs.Todos{Id: 1, Title: "Old Todo", Content: "Old Content", Category_Id: 1, Is_Done: false})

	r := setupRouter(db)

	// 正常な更新
	todoJSON := `{"Title": "Updated Todo", "Content": "Updated Content", "Category_Id": 1, "Is_Done": true}`
	req, _ := http.NewRequest("PUT", "/v1/rest/todo/1", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Todo")

	// 存在しないCategory_Id
	todoJSON = `{"Title": "Updated Todo", "Content": "Updated Content", "Category_Id": 999, "Is_Done": true}`
	req, _ = http.NewRequest("PUT", "/v1/rest/todo/1", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Idに整数以外の値
	todoJSON = `{"Title": "Updated Todo", "Content": "Updated Content", "Category_Id": 1, "Is_Done": true}`
	req, _ = http.NewRequest("PUT", "/v1/rest/todo/abc", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Titleがない場合
	todoJSON = `{"Content": "Updated Content", "Category_Id": 1, "Is_Done": true}`
	req, _ = http.NewRequest("PUT", "/v1/rest/todo/1", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Contentがない場合
	todoJSON = `{"Title": "Updated Todo", "Category_Id": 1, "Is_Done": true}`
	req, _ = http.NewRequest("PUT", "/v1/rest/todo/1", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Is_Doneがない場合
	todoJSON = `{"Title": "Updated Todo", "Content": "Updated Content", "Category_Id": 1}`
	req, _ = http.NewRequest("PUT", "/v1/rest/todo/1", strings.NewReader(todoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Todo")
}

func TestDeleteTodo(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Todos{})

	// Todoを追加
	db.Create(&structs.Todos{Id: 1, Title: "Delete Todo", Content: "Delete Content", Category_Id: 1, Is_Done: false})

	r := setupRouter(db)

	// 正常な削除リクエスト
	req, _ := http.NewRequest("DELETE", "/v1/rest/todo/1", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "消去完了")

	// 存在しないTodoの削除リクエスト
	req, _ = http.NewRequest("DELETE", "/v1/rest/todo/999", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "存在しないTodo")
}
