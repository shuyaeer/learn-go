package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベースの接続情報
	db, err := sql.Open("mysql", "golang:golang@tcp(mysql:3306)/golang")
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer db.Close()

	fmt.Println("Connected to the database!")

	_, err = db.Exec("insert into users (username, email, password, created_at) values (?, ?, ?, now())", "test", "test", "test")
	if err != nil {
		fmt.Println("Failed to insert data:", err)
		return
	}

	// データの取得
	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println("Failed to select data:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var username string
		var email string
		var password string
		var createdAt string
		err := rows.Scan(&id, &username, &email, &password, &createdAt)
		if err != nil {
			fmt.Println("Failed to scan data:", err)
			return
		}
		fmt.Println(id, username, email, password, createdAt)
	}
}

