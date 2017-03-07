package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	class string
}

//结构体可以传指针类型
func (this *Student) Display() {
	fmt.Println(this.Name, ",", this.Age)
}

//定义一个大学生类，继承 Student
type CollegeStudent struct {
	Student
	Profession string
}
func main() {
	s1 := CollegeStudent{Student: Student{Name: "李四 ", Age: 23, class: "2004(2)班 "},Profession: "物理"}
	s1.Display()
	fmt.Println(s1.Student.Name) //可以通过 student 访问 Name
	fmt.Println(s1.Name) //也可以直接通过 name 访问
}
