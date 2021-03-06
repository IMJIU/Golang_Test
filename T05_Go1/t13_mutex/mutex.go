package main

import (
	"errors"
	"fmt"
	"sync"
)

type MyMap struct {
	mp    map[string]int
	mutex *sync.Mutex
}

func (this *MyMap) Get(key string) (int, error) {
	this.mutex.Lock()
	i, ok := this.mp[key]
	this.mutex.Unlock()
	if !ok {
		return i, errors.New("不存在")
	}
	return i, nil
}
func (this *MyMap) Set(key string, v int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.mp[key] = v
}
func (this *MyMap) Display() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	for k, v := range this.mp {
		fmt.Println(k, "=", v)
	}
}
func SetValue(m *MyMap) {
	var a rune
	a = 'a'
	for i := 0; i < 10; i++ {
		m.Set(string(a+rune(i)), i)
	}
}
func main() {
	m := &MyMap{mp: make(map[string]int), mutex: new(sync.Mutex)}
	go SetValue(m) /*启动一个线程向 map 写入值*/
	go m.Display() /*启动一个线程读取 map 的值*/
	var str string /*这里主要是等待线程结束*/
	fmt.Scan(&str)
}
