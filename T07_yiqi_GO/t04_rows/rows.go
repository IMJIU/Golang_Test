package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var row *sql.Rows
	row, err = db.Query("select * from person limit 1")
	if row.Next() {
		//返回所有的列名，列的顺序跟 SCan 一至
		fmt.Println(row.Columns())
		var id, name, age, isBoy interface{}
		err = row.Scan(&id, &name, &age, &isBoy)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("id", id, string(id.([]byte)))
		fmt.Println("name", name, string(name.([]byte)))
	}
}
