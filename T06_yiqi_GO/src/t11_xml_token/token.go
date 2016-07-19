package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Student struct {
	Name string
	Age  int
}

/**
Token 解析是最快的。对于大文件解析，或对性能有要
求时，这种方法是最佳选择。
*/
func main() {
	str := "<?xml version=\"1.0\"encoding=\"utf-8\"?><Student><Name>zhangThree</Name><Age>19</Age></Student>"
	decoder := xml.NewDecoder(strings.NewReader(str))
	var strName string
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			stelm := xml.StartElement(t)
			fmt.Println("Start ", stelm.Name.Local)
			strName = stelm.Name.Local
		case xml.EndElement:
			endelem := xml.EndElement(t)
			fmt.Println("End ", endelem.Name.Local)
		case xml.CharData:
			data := xml.CharData(t)
			str := string(data)
			switch strName {
			case "Name":
				fmt.Println("name:", str)
			case "Age":
				fmt.Println("age:", str)
			default:
				fmt.Println("other:", str)
			}
		}
	}
}
