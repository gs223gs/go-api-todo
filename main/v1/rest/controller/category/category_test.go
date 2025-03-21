package category_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gs223gs/go-api-todo/controller/category"
	"github.com/gs223gs/go-api-todo/structs"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	category.V1RestCategory(r, db)
	return r
}

func testCategoryRequest(t *testing.T, r *gin.Engine, method, path, categoryJSON string, expectedStatus int, expectedBody string) {
	req, _ := http.NewRequest(method, path, strings.NewReader(categoryJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, expectedStatus, w.Code)
	if expectedBody != "" {
		assert.Contains(t, w.Body.String(), expectedBody)
	}
}

func TestGetCategories(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Categories{})
	db.Create(&structs.Categories{Id: 1, Category: "Test Category"})

	r := setupRouter(db)
	testCategoryRequest(t, r, "GET", "/v1/rest/category", "", http.StatusOK, "Test Category")
}

func TestPostCategory(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Categories{})

	r := setupRouter(db)

	tests := []struct {
		name         string
		categoryJSON string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "正常なリクエスト",
			categoryJSON: `{"Category": "New Category"}`,
			expectedCode: http.StatusOK,
			expectedBody: "New Category",
		},
		{
			name:         "不正なTitle",
			categoryJSON: `{"Category": ""}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Category名がありません",
		},
		{
			name:         "不正なJSON",
			categoryJSON: `invalid json`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCategoryRequest(t, r, "POST", "/v1/rest/category", tt.categoryJSON, tt.expectedCode, tt.expectedBody)
		})
	}
}

func TestPutCategory(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Categories{})
	db.Create(&structs.Categories{Id: 1, Category: "Old Category"})

	r := setupRouter(db)

	tests := []struct {
		name         string
		categoryJSON string
		expectedCode int
		expectedBody string
		checkDB      bool
	}{
		{
			name:         "正常な更新",
			categoryJSON: `{"Id": 1, "Category": "Updated Category"}`,
			expectedCode: http.StatusOK,
			expectedBody: "Updated Category",
			checkDB:      true,
		},
		{
			name:         "存在しないID",
			categoryJSON: `{"Id": 999, "Category": "Not Found Category"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "カテゴリがありません",
		},
		{
			name:         "不正なJSON",
			categoryJSON: `invalid json`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
		{
			name:         "不正なTitle",
			categoryJSON: `{"Id": 1, "Category": ""}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Category名がありません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCategoryRequest(t, r, "PUT", "/v1/rest/category", tt.categoryJSON, tt.expectedCode, tt.expectedBody)

			if tt.checkDB {
				var updatedCategory structs.Categories
				if err := db.First(&updatedCategory, 1).Error; err != nil {
					t.Fatalf("更新されたカテゴリが見つかりません: %v", err)
				}
				assert.Equal(t, "Updated Category", updatedCategory.Category)
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&structs.Categories{})
	db.Create(&structs.Categories{Id: 1, Category: "Delete Category"})

	r := setupRouter(db)

	tests := []struct {
		name         string
		categoryJSON string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "正常な削除",
			categoryJSON: `{"Id": 1}`,
			expectedCode: http.StatusOK,
			expectedBody: "",
		},
		{
			name:         "存在しないカテゴリの削除",
			categoryJSON: `{"Id": 999}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: "カテゴリがありません",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCategoryRequest(t, r, "DELETE", "/v1/rest/category", tt.categoryJSON, tt.expectedCode, tt.expectedBody)
		})
	}
}
