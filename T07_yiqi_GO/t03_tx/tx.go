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
	var trans *sql.Tx
	trans, err = db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	trans.Exec("insertintoperson(name,age,IsBoy)values('IMJIU',99,false)")
	if err != nil {
		trans.Rollback()
	} else {
		trans.Commit()
    }
//	trans.Rollback()
}
