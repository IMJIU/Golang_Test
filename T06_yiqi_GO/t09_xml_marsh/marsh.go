package main

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s := &Student{"zhang three", 19}
	result, err := xml.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(result))

	str := `<?xml version="1.0" encoding="utf-8"?>
			<Student>
			<Name>zhang three</Name>
			<Age>19</Age>
			</Student>`
	var s2 Student
	xml.Unmarshal([]byte(str), &s2)
	fmt.Println(s2)
}
