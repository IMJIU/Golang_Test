package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	var mp = map[int]string{1: "a", 2: "b", 3: "c"}
	for k, v := range mp {
		fmt.Println(k, "=", v)
	}
	for _, v := range arr {
		fmt.Println(v)
	}
}
