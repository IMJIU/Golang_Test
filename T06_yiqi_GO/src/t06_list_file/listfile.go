package main

import (
	"fmt"
	"os"
	"io/ioutil"
)
func main(){
	listfile1()
	listfile2()
}
func listfile1() {
	f, err := os.OpenFile("C:\\Windows", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	arrFile, err1 := f.Readdir(0)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	for k, v := range arrFile {
		fmt.Println(k, "\t", v.Name(), "\t", v.IsDir())
	}
}
/**
在 ioutil 包中还提
供了一个 ReadDir 函数，定义如下：
func ReadDir(dirname string) ([]os.FileInfo, error)
读取目录下所有的文件和子目录。是对 File.Readdir 的封装。
*/
func listfile2() {
	arrFile, err := ioutil.ReadDir("C:\\Windows")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for k, v := range arrFile {
		fmt.Println(k, "\t", v.Name(), "\t", v.IsDir())
	}
}

