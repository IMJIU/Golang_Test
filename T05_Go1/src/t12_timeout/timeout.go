package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	select {
	case <-c: //因为没有向 c 发送数据，所以会一直阻塞
		fmt.Print("receive data")
	case <-time.After(5 * time.Second):
		fmt.Println("timeout and quit")
	}
}
