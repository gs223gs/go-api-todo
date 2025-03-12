# 🌟 Welcome to go-webapi-todo

このリポジトリは、よくあるTodoサイト作成に役立つリポジトリです。新しいフレームワークでTodoサイトを作成することはよくあることだと思います。Vue.js, React, Angular... フロントエンドフレームワークは目まぐるしく進化しています。

最初のアウトプットでTodoサイトを開発するのはよくある話です。バックエンドなしでもできますが、あった方が学びになります。しかし、その度にバックエンドを用意するのは面倒ですよね。管理も面倒です。それをこのリポジトリで解決します。

> **⚠️ 注意:** まだ開発途中で使用することはできません。もうしばらくお待ちください。

## 🚀 環境構築

1. **リポジトリをクローン**
   ```bash
   git clone https://github.com/gs223gs/go-api-todo
   ```

2. **ディレクトリに移動**
   ```bash
   cd go-webapi-todo
   ```

3. **コンテナを起動**
   ```bash
   docker compose up -d
   docker compose exec go bash
   ```

4. **`go.mod`を作成**
   ```bash
   go mod init github.com/gs223gs/go-webapi-todo
   ```

5. **APIを起動**
   ```bash
   cd 使いたいAPIに移動 例 => v1/rest
   air
   ```

6. **お好きなAPIにアクセス**
   - JSONでやりとりします。

## 🔮 今後実装する機能

- ユーザー認証 OAuth
- ログ
- RESTful API
- gRPC

---

このプロジェクトに興味を持っていただき、ありがとうございます！フィードバックや貢献をお待ちしております。😊


