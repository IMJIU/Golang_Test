package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s := Student{Name: "abc", Age: 19}
//	rv := reflect.ValueOf(s)

	//这里传的是&s 因为要修改字段的地址，否在会报错。
	rv := reflect.ValueOf(&s)
	
	//判断是否指针类型，如果是，取指针所指向的元素的类型
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	rvField := rv.FieldByName("Name") //取 Name 字段的值
	fmt.Println(rvField.String())
	
	rvField.SetString("change name!")
	fmt.Println(s.Name) //输出已改名
}
