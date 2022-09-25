package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 官方文档 https://pkg.go.dev/github.com/go-sql-driver/mysql#section-readme
// 第一步：设置mysql驱动
func main() {
	// 获取数据库连接
	db, err := sql.Open("mysql", "root:root@/demo")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("连接成功")
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("开始插入数据")

	insert(db)

	fmt.Println("开始查询单行")

	queryOneRow(db, 1)

	fmt.Println("开始查询多行")

	queryList(db)

	fmt.Println("开始更新数据")

	update(db)

	fmt.Println("更新万 查询多行数据")

	queryList(db)

	fmt.Println("开始删除数据")

	delete(db)

	fmt.Println("删除后 查询多行数据")

	queryList(db)

}

// mysql插入数据
func insert(db *sql.DB) {
	s := "insert into user (name) value (?) "
	_, err := db.Exec(s, "wang")
	if err != nil {
		fmt.Println("-------")
		fmt.Println(err)
	}
}

type user struct {
	id   int
	name string
}

// 查询单行
func queryOneRow(db *sql.DB, id int) {
	s := "select * from user where id = ?"
	var u user
	err := db.QueryRow(s, id).Scan(&u.id, &u.name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
}

// 查询多行
func queryList(db *sql.DB) {

	s := "select * from user"
	var u user
	list, err := db.Query(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
	for list.Next() {
		list.Scan(&u.id, &u.name)
		fmt.Println(u)
	}
}

// 更新
func update(db *sql.DB) {
	s := "update user set name = ? where name = ?"
	_, err := db.Exec(s, "jinwei", "wang")
	if err != nil {
		fmt.Println(err)
	}
}

// 删除
func delete(db *sql.DB) {
	s := "delete from user where name = ?"
	_, err := db.Exec(s, "jinwei")
	if err != nil {
		fmt.Println(err)
	}

}
