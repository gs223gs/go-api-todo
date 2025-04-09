package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// ? Todo関係
// !---------------------------------------------------------------------------------
type Todos struct {
	Todo []Todo `json:"todos"`
}

type Todo struct {
	Id     int    `json:"id" gorm:"primary_key;auto_increment"`
	Title  string `json:"title" gorm:"size:255"`
	IsDone bool   `json:"isDone" gorm:"default:false"`
}

func (t Todo) CheckTitle() error {
	return validateString(t.Title, "TodoTitle")
}

func (t Todo) CheckId() error {
	if t.Id <= 0 {
		return fmt.Errorf("Todo IDがないか、0以下です")
	}
	return nil
}

// !---------------------------------------------------------------------------------

// ? validate
// !---------------------------------------------------------------------------------

// (Title or UserName, Type)
func validateString(s, t string) error {
	if s == "" || len(s) > 255 {
		return fmt.Errorf("%sがないか、255文字以上です", t)
	}
	return nil
}

// !---------------------------------------------------------------------------------

// ? error
// !---------------------------------------------------------------------------------

type ErrResponse struct {
	Error []ErrMessage `json:"errors"`
}

type ErrMessage struct {
	Tittle string `json:"tittle"`
	Detail string `json:"detail"`
}

//!---------------------------------------------------------------------------------

// ? userregister
// !---------------------------------------------------------------------------------
type User struct {
	Id       int    `json:"id" gorm:"primary_key;auto_increment"`
	UserName string `json:"userName" gorm:"size:255"`
	Password string `json:"password" gorm:"size:255"`
}

func (u User) CheckUserName() error {
	return validateString(u.UserName, "UserName")
}

func (u User) CheckPassword() error {
	return validateString(u.Password, "Password")
}

//!---------------------------------------------------------------------------------

func CreateJWT() string {
	// 環境変数から秘密鍵を取得
	secretKey := os.Getenv("ACCESS_SECRET_KEY")
	

	claims := jwt.MapClaims{
		"user_id": "user_id1234",
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 72時間が有効期限
	}

	// ヘッダーとペイロード生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("トークン生成エラー: %v", err)
		return ""
	}

	return accessToken
}

func main() {
	UpdateRequest := Todo{0, "", false}

	if err := UpdateRequest.CheckTitle(); err != nil {
		fmt.Println(err)
	}

	if err := UpdateRequest.CheckTitle(); err != nil {
		fmt.Println(err)
	}

	if err := UpdateRequest.CheckId(); err != nil {
		fmt.Println(err)
	}

	response := Todo{1, "test", false}
	todolist := Todos{Todo: []Todo{response, response}}

	jsonTodolist, err := json.Marshal(todolist)
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return
	}
	fmt.Println(string(jsonTodolist))

	errmsg := ErrMessage{"text", "aa"}
	errResponse := ErrResponse{Error: []ErrMessage{errmsg, errmsg, errmsg}}

	jsonResponse, err := json.Marshal(errResponse)
	if err != nil {
		fmt.Println("JSON変換エラー:", err)
		return
	}
	fmt.Println(string(jsonResponse))

}

/*
Todo
DB-struct
gorm
gin
JWT
api-end-point{
	/login[
		POST
	]

	/signup[
		POST
	]

	/todo[
		GET
		POST
		PUT
		DELETE
	]
}


*/
