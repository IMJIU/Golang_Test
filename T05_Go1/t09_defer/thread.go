package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer func() {
		fmt.Println("in defer!")
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i > 5 {
			runtime.Goexit()
		}
	}
}
func main() {
	go test()
	var str string
	fmt.Scan(&str)
}
