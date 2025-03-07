# 今後実装したい機能
- RESTfull API
  - ページング処理
- 認証認可
  - google
  - github
- gRPC
## REST
### GET v1/rest/todo
機能:Todoを検索

レスポンス
```json
{
    "Id": int,
    "Title": string,
    "Content": string,
    "Category_id": int,
    "Is_Done": bool,
    "Due": string("") || time (UTC),
    "Created_at": time (UTC)
}
```

### POST v1/rest/todo
機能:Todoの登録
期待するJSON
```JSON
{
  "Title": string,
  "Content": string,
  "category_id": bigint
}
```

レスポンス
```JSON
{
    "Id": bigint,
    "Title": string,
    "Content": string,
    "Category": string,
    "Is_Done": bool,
    "Created_at": time (ISO 8601 UTC時間),
    "Update_at": time (ISO 8601 UTC時間)
}
```



Contentについて:
  - todoの詳細な説明 null 可能

Category_idについて:
- category tableに
### PUT v1/rest/todo
機能:
### DELETE v1/rest/todo
機能:


### GET v1/rest/category
機能:
### POST v1/rest/category
機能:
### PUT v1/rest/category
機能:
### DELETE v1/rest/category