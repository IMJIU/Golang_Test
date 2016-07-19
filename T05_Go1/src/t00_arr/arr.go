package main

import (
	"fmt"
)

func main() {
	
	//声明一个2个元素的数组，名字为 arr_1,因为是 int 型数组，所以初值为0，即[0,0]
	var arr_1 [2]int
	
	/*声明一个2个元素的数组，名字为 arr_2，并同时赋初值，{}里为空，说明没有赋初
	  值，等同于上面*/
	arr_2 := [2]int{}
	
	//声明一个2个元素的数组，名字为 arr3, arr3_1, arr3_2，并同时赋初值，结果均为[1,2]
	arr3 := [2]int{1, 2}
	
	//{}里的冒号左边是下标，右边是值
	arr3_1 := [2]int{0: 1, 1: 2}
	arr3_2 := [2]int{1: 2, 0: 1}
	
	/*不指定数组长度，自动计算长度, [...],声明一个2个（自动计算而来）元素的数组，名字
	  为 arr4，并同时赋初值，结果为[1,2]*/
	arr4 := [...]int{1, 2}
	
	/*声明一个4个（自动计算而来）元素的数组，名字为 shuzu5，并同时赋初值，结果
	  为[0,0,0,9]*/
	arr5 := [...]int{3: 9}
	
	fmt.Println(arr_1)
	fmt.Println(arr_2)
	fmt.Println(arr3)
	fmt.Println(arr3_1)
	fmt.Println(arr3_2)
	fmt.Println(arr4)
	fmt.Println(arr5)
}
