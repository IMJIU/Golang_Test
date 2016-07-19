package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	err := ioutil.WriteFile("c:\\a.txt", []byte("abcdefg"), 0777)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("OK")
	}
}
