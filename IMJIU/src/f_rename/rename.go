package main

import (
	"os"
	"fmt"
//	"path/filepath"
	"io"
	"strings"
)


func main(){
	f,err := os.OpenFile("g:/book_name.sql",os.O_RDONLY,0666)
	if err!=nil{
		fmt.Println("error!"+err.Error())
		os.Exit(0)
	}
	defer f.Close()
	//读取文件文本
	str:=read(f)
	
	//解析成map
	//fmt.Println(str)
	kv:=parse(str)
	fmt.Println(kv)
	
	//读取需要更改名称的文件
	rename(kv)
}
/**重命名*/
func rename(kv map[string]string){
	f2, err2 := os.OpenFile("g:/book", os.O_RDONLY, 0666)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	defer f2.Close()
	arrFile, err1 := f2.Readdir(0)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	for _, v := range arrFile {
		if !v.IsDir(){
			//fmt.Println( v.Name(), "\t", v.IsDir())
			idx:=strings.Index(v.Name(),".")
			key:=v.Name()[0:idx]
			fname, ok :=kv[key]
			if ok{
			  fmt.Println("g:/book/"+v.Name())
			  fmt.Println(os.Rename("g:/book/"+v.Name(),"g:/book/"+fname))
				

//			  fmt.Println("success!")

			}
		}
		//fmt.Println(k, "\t", v.Name(), "\t", v.IsDir())
		//fmt.Println(filepath.Abs(v.Name()))
		//fmt.Println(filepath.Abs("/"))
	}
}
/** 解析成map */
func parse(s string)map[string]string{
	arr := strings.Split(s, "\n")
	//var ms []map[string]string  = make([]map[string]string,len(arr))
	m := make(map[string]string)
	for _,v := range arr{
		//fmt.Println(k,v)
		vs := strings.Split(v,"-")
		if vs[0][0]=='[' {
			vs[0]=strings.Trim(strings.Replace(vs[0],"[","",-1)," ");
			vs[0]=strings.Trim(strings.Replace(vs[0],"]","",-1)," ");
			if len(vs)>2{
				m[vs[0]]=strings.Trim(vs[1]," ")+strings.Trim(vs[2]," ");
			}else{
				m[vs[0]]=strings.Trim(vs[1]," ");
			}
		}
	}
	return m
}
/** 读取文本内容*/
func read(f *os.File)string{
	b:=make([]byte,1024)
	f.Seek(0,os.SEEK_SET)
	str:=""
	for{
		len,ferr:=f.Read(b)
		if ferr!=nil&&ferr!=io.EOF{
			fmt.Println(ferr.Error())
			break
		}
		if len==0{
			break
		}
		//fmt.Println(len)
		str += string(b[0:len])
	}
	return str
}