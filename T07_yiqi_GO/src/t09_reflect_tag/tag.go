package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "studentName"
	Age  int    `a:"1111"b:"3333"` //这个不是单引号，而是~键上的符号
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")
	//取 tag 数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")
	/*  可以你 JSON 一样，取 TAG 里的数据，注意，设置时，两个之间无逗 号,键名无引号  */
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}
}
