package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	student := &Student{"zhang three", 19}
	f, err := os.Create("data.dat")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	//创建 Encoder 对像
	encode := gob.NewEncoder(f)
	//将 s 序列化到 f 文件中
	encode.Encode(student)
	//重置文件指针到开始位置
	f.Seek(0, os.SEEK_SET)
	decoder := gob.NewDecoder(f)
	var s1 Student
	//反序列化对像
	decoder.Decode(&s1)
	fmt.Println(s1)
}
