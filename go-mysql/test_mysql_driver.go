package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// https://pkg.go.dev/github.com/go-sql-driver/mysql#section-readme
// 第一步：设置mysql驱动
func main() {

	db, err := sql.Open("mysql", "root:root@/local")
	if err != nil {
		fmt.Println(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

}
