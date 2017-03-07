package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	class string
}

func (this Student) getName() string {
	return this.Name
}

//结构体可以传指针类型
func (this *Student) getAge() int {
	return this.Age
}

func main() {

	//方式1：
	s1 := new(Student)
	s1.Name = "张三"
	s1.Age = 12
	s1.class = "21班"
	fmt.Println(s1)

	//方式2：
	s1 := Student{Name: "张三", Age: 12, class: "2 班"}

	//方式3：
	s1 := Student{"张三", 12, "2 班"}
	
	fmt.Println(s.getName(), s.getAge())
}
