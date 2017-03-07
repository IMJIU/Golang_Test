package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var smt *sql.Stmt
	smt, err = db.Prepare("insert into person(name,age,IsBoy)values(?,?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("begin....", time.Now())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		_, err = smt.Exec(fmt.Sprintf("zhang%d", r.Int()), r.Intn(50),
			r.Intn(100)%2)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("success!", time.Now())
}
