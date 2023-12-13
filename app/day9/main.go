package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/shuyaeer/learn-go/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
  // データベースへの接続設定
  db, err := sql.Open("mysql", "golang:golang@tcp(mysql:3306)/golang?parseTime=true")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // データベース接続のテスト
  if err = db.Ping(); err != nil {
      log.Fatal(err)
  }

  // sqlboiler用の設定
  boil.SetDB(db)

  // usersテーブルから全ユーザーを取得
  users, err := models.Users().All(context.Background(), db)
  if err != nil {
      log.Fatal(err)
  }

  // 取得したユーザー情報を表示
  for _, user := range users {
      fmt.Printf("User: %+v\n", user)
  }
}
