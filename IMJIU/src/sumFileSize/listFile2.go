 package main
 
 import (
     //"flag"
     "fmt"
     "os"
     "path/filepath"
     "time"
//     "strings"
 )
 
 var (
     ostype = os.Getenv("GOOS") // 获取系统类型
     fcount int64 = 0
 )
 
 var listfile []string //获取文件列表
 
 func Listfunc(path string, f os.FileInfo, err error) error {
     var strRet string
     strRet, _ = os.Getwd()
     //ostype := os.Getenv("GOOS") // windows, linux
     if ostype == "windows" {
         strRet += "\\"
     } else if ostype == "linux" {
         strRet += "/"
     }
     if f == nil {
         return err
     }
     if f.IsDir() {
         return nil
     }
     //strRet += path //+ "\r\n"
     fcount += f.Size()
     //用strings.HasSuffix(src, suffix)//判断src中是否包含 suffix结尾
//     ok := strings.HasSuffix(strRet, ".go")
//     if ok {
//         listfile = append(listfile, strRet) //将目录push到listfile []string中
//     }
     //fmt.Println(ostype) // print ostype
     //fmt.Println(strRet) //list the file
     return nil
 }
 
 func getFileList(path string) string {
     //var strRet string
     err := filepath.Walk(path, Listfunc) //
     if err != nil {
         fmt.Printf("filepath.Walk() returned %v\n", err)
     }
     return " "
 }
 
 func ListFileFunc(p []string) {
     for index, value := range p {
         fmt.Println("Index = ", index, "Value = ", value)
     }
 }
 
 func main() {
     //flag.Parse()
     //root := flag.Arg(0)
     //fmt.Println()
     
     var listpath string="f:/"
     //fmt.Scanf("%s", &listpath)
     begin := time.Now().Second()
     getFileList(listpath)
     //ListFileFunc(listfile)
     fmt.Println("耗时：",time.Now().Second()-begin,"秒")
     fmt.Println("总共：",fcount/1024/1024/1024,"G")
     //getFileList(root)
 
 }