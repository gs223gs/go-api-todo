package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// JWTに付与するクレームを定義
	claims := jwt.MapClaims{
		"user_id": "user_id123", // ユーザーIDをクレームに追加
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // トークンの有効期限を72時間後に設定
	}

	// JWTトークンを生成し、ヘッダーとペイロードを設定
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与し、署名付きトークンを生成
	accessToken, _ := token.SignedString([]byte("ACCESS_SECRET_KEY"))
	fmt.Println("accessToken:", accessToken) // 生成されたアクセストークンを出力

	// 生成されたトークン文字列を変数に格納
	tokenString := accessToken

	// トークンを解析し、署名の検証を行う
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名方法が期待通りか確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"]) // 予期しない署名方法の場合のエラーメッセージ
		}

		// 署名検証のためのキーを返す
		return []byte("ACCESS_SECRET_KEY"), nil
	})

	// トークンのクレームを取得し、トークンが有効か確認
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("user_id: %v\n", string(claims["user_id"].(string))) // クレームからユーザーIDを取得し出力
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64))) // クレームから有効期限を取得し出力
	} else {
		fmt.Println(err) // トークンが無効な場合のエラーメッセージを出力
	}
}
