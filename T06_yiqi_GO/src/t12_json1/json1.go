package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age  int
}
type Student2 struct {
	Name string `json:"userName"`
	Age  int
}
func main(){
	json1()
	json2()
}
func json1() {
	f, err := os.Create("data.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := &Student{"zhangThree", 19}
	//创建 encode 对像
	encoder := json.NewEncoder(f)
	//将 s 序列化到文件中
	encoder.Encode(s)
	//重置文件指针到开始位置
	f.Seek(0, os.SEEK_SET)
	decoder := json.NewDecoder(f)
	var s1 Student
	//从文件中反序列化成对像
	decoder.Decode(&s1)
	fmt.Println(s1)
}


func json2() {
	s := &Student2{"zhangThree", 19}
	//将 s 编码为 json
	buf, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(buf))
	//将 json 字符串转换成 Student 对像
	var s1 Student2
	json.Unmarshal(buf, &s1)
	fmt.Println(s1)
}

