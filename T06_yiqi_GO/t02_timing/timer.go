package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	c := time.After(10 * time.Second) //返回 channel 类型,10秒后向 channel 发送当前时间
	t := <-c
	fmt.Println(t)
	tm := time.NewTimer(10 * time.Second) //NewTimer 返回 Timer 类型
	t = <-tm.C                            /*Timer 结构中有一个 channel C,10秒后，把当前时间发送到 C*/
	fmt.Println(t)

	fmt.Println("====== AfterFunc() ==============")
	f2 := func() {
		fmt.Println("Hello world!", time.Now())
	}
	fmt.Println(time.Now())
	time.AfterFunc(2 * time.Second, f2)

	//	var str string
	//	fmt.Scan(&str) /*这里主要是等待用户输入，不让进程结束，进程结束定时器也就无效了。*/

	fmt.Println("====== Internal() ==============")
	c2 := time.Tick(1 * time.Second)
	for t := range c2 {
		fmt.Println(t)
	}
}
