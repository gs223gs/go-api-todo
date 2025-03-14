package validation_test

import (
	"testing"

	"gorm.io/gorm"
	"github.com/gs223gs/go-webapi-todo/controller/validation"
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
