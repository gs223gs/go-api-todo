# Todo API仕様書


## エンドポイント

### Todoリソース

#### 1. Todo一覧取得
**エンドポイント**: `GET /v1/rest/todo`

**レスポンス**:
```json
[
  {
    "Id": 1,
    "Title": "タスク名",
    "Content": "タスク内容",
    "Category_id": 1,
    "Is_Done": false, //defaultはfalse
    "Due": "2024-03-21T13:36:25Z",  // null の場合は空文字
    "Created_at": "2024-03-21T13:36:25Z",
    "Updated_at": "2024-03-21T13:36:25Z"
  }
]
```

#### 2. Todo作成
**エンドポイント**: `POST /v1/rest/todo`

**リクエストボディ**:
```json
{
  "Title": "タスク名",
  "Content": "タスク内容",
  "Category_Id": 1,
  "Is_Done": false, //なくてもいい defaultは false
  "Due": "2024-03-21T13:36:25Z"  // なくてもいい defaultは空文字
}
```

**レスポンス**:
```json
{
  "messege": "追加完了"
}
```

#### 3. Todo更新
**エンドポイント**: `PUT /v1/rest/todo`

**リクエストボディ**:
```json
{
  "Id": 1,
  "Title": "更新後のタスク名",
  "Content": "更新後のタスク内容",
  "Category_Id": 1,
  "Is_Done": true,
  "Due": "2024-03-21T13:36:25Z"  // オプション
}
```

**レスポンス**:
```json
{
  "message": "更新完了"
}
```

#### 4. Todo削除
**エンドポイント**: `DELETE /v1/rest/todo`

**リクエストボディ**:
```json
{
  "Id": 1
}
```

**レスポンス**:
```json
{
  "messege": "消去完了"
}
```

### カテゴリーリソース

## エラーレスポンス
```json
{
  "error": "エラーメッセージ"
}
```

## バリデーション
- Todo名は必須
- カテゴリー名は必須
- カテゴリーIDは存在するものを指定
- TodoIDは存在するものを指定

## 注意事項
- カテゴリーを削除すると、関連するTodoも削除されます（カスケード削除）
- 日時はISO 8601形式で扱います
- すべてのエンドポイントは JSON 形式でデータをやり取りします