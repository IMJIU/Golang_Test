package main

import (
	"os"
	"fmt"
	"path"
	"strings"
	"path/filepath"
)

func main() {
	fmt.Println(path.Base("/usr/bin"))    //输了出 bin
	fmt.Println(path.Base(""))            //输出. 
	fmt.Println(path.Base("C:\\Windows")) /*无法识别 Windows 下的
	路径分隔符，将会把 C:\\Windows 做为一个路径*/
	fmt.Println(path.Base(strings.Replace("C:\\Windows", "\\", "/", -1))) /*把\转换成/*/
	
	fmt.Println("======== Clean ============");
	fmt.Println(path.Clean("/a/b/../c")) /*/a/c*/
	fmt.Println(path.Clean("/a/b/../././c")) /*/a/c*/
	
	fmt.Println("======== dir ============");
	fmt.Println(path.Dir("/a/b/../c/d/e")) /*/a/c/d*/
	fmt.Println(path.Clean("/a/b/")) /*/a/b*/
	
	//扩展名
	fmt.Println("======== Ext ============");
	fmt.Println(path.Ext("/a/b/../c/d./e")) /*没有扩展名*/
	fmt.Println(path.Ext("/a/b/test.txt")) /*.txt*/
	
	//是否绝对路径
	fmt.Println("======== IsAbs ============");
	fmt.Println(path.IsAbs("/a/b/c"))
	fmt.Println(path.IsAbs(strings.Replace("C:\\Windows\\system","\\", "/", -1))) /*Go 只识别/所以需要转换一下*/
	
	fmt.Println("======== Join ============");
	fmt.Println(path.Join("a/b", "c"))/*a/b/c*/
	fmt.Println(path.Join("C:\\Windows","System"))/*C:\Windows/System*/
	
	fmt.Println("======== Split============");
	fmt.Println(path.Split("/a/b/test.txt")) /*/a/b/ test.txt*/
	fmt.Println(path.Split("/a/b/c/")) /*/a/b/c/ */
	
	fmt.Println("======== Abs ============");
	fmt.Println(filepath.Abs("."))
	
	fmt.Println("======== Walk ============");
	filepath.Walk(".", DispFile)
	
	
}
func DispFile(path string, info os.FileInfo, err error) error {
	fmt.Println(path,"-------",info.Name(), "------", info.IsDir())
	return nil
}
