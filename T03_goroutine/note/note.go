package t02

import "sync"

func none(){
	var ch1 chan int // ch1是一个正常的channel，不是单向的
	var ch2 chan<- float64// ch2是单向channel，只用于写float64数据
	var ch3 <-chan int // ch3是单向channel，只用于读取int数据
	
	ch4 := make(chan int)
	ch5 := <-chan int(ch4) // ch5就是一个单向的读取channel
	ch6 := chan<- int(ch4) // ch6 是一个单向的写入channel
	
	close(ch)
	//在介绍了如何关闭channel之后，我们就多了一个问题：如何判断一个channel是否已经被关
	//闭？我们可以在读取的时候使用多重返回值的方式：
	x, ok := <-ch
	//这个用法与map中的按键获取value的过程比较类似，只需要看第二个bool返回值即可，如
	//果返回值是false则表示ch已经被关闭
	
	
	var l sync.Mutex
	func foo() {
		l.Lock()
		defer l.Unlock()
		//...
	}
}
