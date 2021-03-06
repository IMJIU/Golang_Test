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
	rt := reflect.TypeOf(s)
	//判断是否指针类型，如果是，取指针所指向的元素的类型
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	//输出类型所在的包的路径
	fmt.Println(rt.PkgPath())
	//反射取所有字段
	fmt.Println(rt.Name(), "total:", rt.NumField(), " property")
	for i, j := 0, rt.NumField(); i < j; i++ {
		rtField := rt.Field(i)
		fmt.Println(rtField.Name)
	}
	/*因为我们的函数定义是在*Student 类型上的，
	所以这里转换为指针类型， 否  则反射会取不到函数*/
	rt = reflect.PtrTo(rt)
	//反射取所有函数
	fmt.Println(rt.Name(), "total:", rt.NumMethod(), " func")
	for i, j := 0, rt.NumMethod(); i < j; i++ {
		mt := rt.Method(i)
		fmt.Println(mt.Name)
		//输入参数的数量
		numIn := mt.Type.NumIn()
		//输出参数的数量
		numOut := mt.Type.NumOut()
		//输出输入参数
		if numIn > 0 {
			fmt.Println("\t total:", numIn, " in param")
			for k := 0; k < numIn; k++ {
				in := mt.Type.In(k)
				fmt.Println("\t", in.Name(), "\t", in.Kind())
			}
		}
		//输出输出参数
		if numOut > 0 {
			fmt.Println("\t total:", numOut, "out param")
			for k := 0; k < numOut; k++ {
				out := mt.Type.Out(k)
				fmt.Println("\t", out.Name(), "\t", out.Kind())
			}
		}
	}
}
/**
main
Student total: 2  property
Name
Age
 total: 2  func
GetAge
	 total: 1  in param
	  	 ptr
	 total: 1 out param
	 int 	 int
PrintName
	 total: 1  in param
	  	 ptr
*/
