package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (this *Student) PrintName() {
	fmt.Println(this.Name)
}
func (this *Student) GetAge() int {
	return this.Age
}
func main() {
	s := Student{Name: "abc", Age: 19}
	rt := reflect.TypeOf(&s)
	rv := reflect.ValueOf(&s)
	fmt.Println("typeof func,PrintName:")
	rtm, ok := rt.MethodByName("PrintName")
	if ok {
		var parm []reflect.Value
		//函数默认第一个参数是结构体本身即*Student
		parm = append(parm, rv)
		rtm.Func.Call(parm)
	}
	//valueof 调用函数
	fmt.Println("valueof func:")
	rvm := rv.MethodByName("GetAge")
	//用 valueof 调用函数时不需要把 Struct 本身做为参数传递过去
	ret := rvm.Call(nil)
	//显示返回值
	fmt.Println("return value:")
	ShowSlice(ret)
}
func ShowSlice(s []reflect.Value) {
	if s != nil && len(s) > 0 {
		for _, v := range s {
			fmt.Println(v.Interface())
		}
	}
}
