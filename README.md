# 🌟 Welcome to go-webapi-todo

このリポジトリは、よくあるTodoサイト作成に役立つリポジトリです。新しいフレームワークでTodoサイトを作成することはよくあることだと思います。Vue.js, React, Angular... フロントエンドフレームワークは目まぐるしく進化しています。

最初のアウトプットでTodoサイトを開発するのはよくある話です。バックエンドなしでもできますが、あった方が学びになります。しかし、その度にバックエンドを用意するのは面倒ですよね。管理も面倒です。それをこのリポジトリで解決します。

> **⚠️ 注意:** v1/rest 使用可能となりました．

## 🚀 環境構築

1. **リポジトリをクローン**
   ```bash
   git clone https://github.com/gs223gs/go-api-todo
   ```

2. **ディレクトリに移動**
   ```bash
   cd go-api-todo
   ```

3. **コンテナを起動**
   ```bash
   docker compose up -d
   docker compose exec go bash
   ```

4. **`go.mod`を作成**
   ```bash
   cd 使いたいAPIに移動 => 例: v1/rest
   go mod init github.com/gs223gs/go-api-todo
   go mod tidy
   ```

5. **APIを起動**
   ```bash
   go run main.go
   ```

6. **お好きなAPIにアクセス**
   - JSONでやりとりします。

## 余談
VScodeの拡張機能を使用して開発コンテナ内に入って作業していました
もし，goが使用できない場合そちらで解決するかもしれません．
## 🔮 今後実装する機能

- ~~RESTful API~~ 2025/03/22 完成 v1/rest
- OAuth
- Open ID connection
- JWT
- ログ出力
- GraphQL
- gRPC
- 並行処理を用いたもの
   - Todo登録のたびにメール送信?
   - 外部APIを呼び出す何かしらの機能？

Qiita

私が学んだことや，改善点として気がついたことについては以下をご覧ください

[Qiita 私が気がついたこと](https://qiita.com/gs223gs/items/402893d1194c9ef26d7c)

このプロジェクトに興味を持っていただき、ありがとうございます！フィードバックや貢献をお待ちしております。😊


