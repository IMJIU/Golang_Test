package main

import (
	"fmt"
)

//MyFuncType 是一个接受 int 类型的参数，返回 bool 值的函数类型
type MyFuncType func(int) bool

func IsBigThan5(n int) bool {
	return n > 5
}

func Display(arr []int, f MyFuncType) {
	for _, v := range arr {
		if f(v) {
			fmt.Println(v)
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 0, 0, -1, 9}
	Display(arr, IsBigThan5)
}
