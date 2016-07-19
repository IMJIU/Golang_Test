package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age  int
	getName func() string
}

func main() {
	f, err := os.Create("data.xml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	f2 := func()string{
		fmt.Println("f2()")
		return "kjdfa"
	};
	s := &Student{"zhang three ", 19,f2}
	//创建 encode 对像
	encoder := xml.NewEncoder(f)
	//将 s 序列化到文件中
	encoder.Encode(s)
	//重置文件指针到开始位置
	f.Seek(0, os.SEEK_SET)
	decoder := xml.NewDecoder(f)
	var s1 Student
	//从文件中反序列化成对像
	decoder.Decode(&s1)
	fmt.Println(s1)
}
