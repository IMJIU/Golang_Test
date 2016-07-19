package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
/**
DROP TABLE IF EXISTS `person`;
CREATE TABLE `person` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`name` varchar(255) DEFAULT NULL,
`age` int(11) DEFAULT NULL,
`IsBoy` tinyint(4) DEFAULT NULL,
PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;
*/
func main() {
	db, err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var result sql.Result
	//向数据库中插入一条数据
	result, err = db.Exec("insert into person(name,age,IsBoy)values(?,?,?)", "zhangThree", 19, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	lastId, _ := result.LastInsertId()
	fmt.Println("lastId:", lastId)
	var row *sql.Row
	//返回一行数据
	row = db.QueryRow("select * from person")
	var name string
	var id, age int
	var isBoy bool
	//取数据进行显示
	err = row.Scan(&id, &name, &age, &isBoy)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id, "\t", name, "\t", age, "\t", isBoy)
	//再插入一条数据
	result, err = db.Exec("insert into person(name,age,IsBoy) values(?,?,?)", "wangHong", 18, false)
	fmt.Println("=====================")
	var rows *sql.Rows
	rows, err = db.Query("select * from person")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var name string
		var id, age int
		var isBoy bool
		rows.Scan(&id, &name, &age, &isBoy)
		fmt.Println(id, "\t", name, "\t", age, "\t", isBoy)
	}
	rows.Close()
	// 最后，清空表
	db.Exec("truncate table person")
}
