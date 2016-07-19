package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main() {
//	readfile()
//	readall()
	readfile2()
}
func readfile(){
	/*打开 D:\\新建文本文档.txt 文件，如果文件不存在将会新建，如果已
	存在，新写入的内容将追加到文件尾*/
	f, err := os.OpenFile("c:/atlog.txt", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	f.WriteString("\r\n 中国\r\n")
	buf := make([]byte, 1024)
	var str string
	/*重置文件指针，否则读不到内容的。*/
	f.Seek(0, os.SEEK_SET)
	for {
		n, ferr := f.Read(buf)
		if ferr != nil && ferr != io.EOF {
			fmt.Println(ferr.Error())
			break
		}
		if n == 0 {
			break
		}
		fmt.Println(n)
		str += string(buf[0:n])
	}
	fmt.Println(str)
	f.Close()
}
func readall(){
	f, err := os.OpenFile("c:\\atlog.txt",os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	buf, err1 := ioutil.ReadAll(f)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	fmt.Println(string(buf))
}
func readfile2() {
	buf, err := ioutil.ReadFile("c:\\atlog.txt")
	if err != nil {
	fmt.Println(err.Error())
	return
	}
	fmt.Println(string(buf))
}