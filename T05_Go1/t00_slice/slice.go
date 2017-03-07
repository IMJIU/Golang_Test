package main

import "fmt"
import "reflect"

func main() {
	p := [...]int{2, 3, 5, 7, 11, 13} //定义一个数组
	s1 := p[1:3]                      //定义切片，包含3,5两个元素
	
	fmt.Println(s1)
	fmt.Println(reflect.TypeOf(p))  //用反射得到变量的类型 p 是数组类型[6]int
	fmt.Println(reflect.TypeOf(s1)) //s1是切片类型[]int
	
	ChangeArrayValue(p)             //ChangeArrayValue 函数将第一个值改为100
	fmt.Println(p)                  //数组 p 的值并没有改变，因为数组是值类型
	
	ChangeSliceValue(s1)            //ChangeSliceValue 将切片的第一个值改为100
	fmt.Println(s1)                 //切片 s1的值被改变，因为切片是引用类型
	fmt.Println(p)                  //切片是引用的数组 p 第一个和第二个元素，所以数组 p 的值被改变
}

func ChangeArrayValue(arr [6]int) {
	arr[0] = 100
}

func ChangeSliceValue(slice []int) {
	slice[0] = 100
}
