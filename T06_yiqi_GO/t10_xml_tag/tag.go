package main

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	XMLName xml.Name `xml:"student"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

type Student2 struct {
	XMLName xml.Name `xml:"student"`
	Name    string   `xml:"name,attr"`
	Age     int      `xml:"age,attr"`
	Phone   []string `xml:"phones>phone",`
}

type ABC string

func main(){
	parse1()
	parse2()
}

func parse1() {
	str := `<?xml version="1.0" encoding="utf-8"?>
			<student>
			<name>zhang three</name>
			<age>19</age>
			</student>`
	var s Student
	xml.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}



func parse2() {
	str := `<?xml version="1.0" encoding="utf-8"?>
			<student name="zhang three" age="19">
			<phones>
			<phone>12345</phone>
			<phone>67890</phone>
			</phones>
			</student>`
	var s Student2
	xml.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}