package main

import (
	"fmt"
)

func main() {
	mp := make(map[string]string) //key 是字符串类型，值也是字符串类型
	mp["a"] = "1"
	mp["b"] = "2"
	mp["pi"] = "3.1415926"
	mp["sh"] = "上海"
	v, ok := mp["sh"] //sh 存在，v 存放的是 value 值，ok 值为 true
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key 'sh'' 不存在")
	}
	v, ok = mp["bj"] //bj 不存在，ok 为 false
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key 'bj' 不存在")
	}
}
