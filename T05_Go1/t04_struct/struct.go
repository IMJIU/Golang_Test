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

func (this Student) change()  {
	 this.Age=5555
}
func (this *Student) change2()  {
	this.Age=5555
}

func main() {

	//方式1：
	 s1 := new(Student)
	s1.Name = "张三"
	s1.Age = 12
	s1.class = "21班"
	fmt.Println(s1)

	//方式2：
	s2 := Student{Name: "张三", Age: 12, class: "2 班"}

	//方式3：
	s3 := Student{"张三", 12, "2 班"}
	
	fmt.Println(s1.getName(), s2.getAge(),s3.class)
	fmt.Println("==============")
	s1.change()
	fmt.Println(s1.getAge())
	s1.change2()//引用类型才能改变值
	fmt.Println(s1.getAge())
}
