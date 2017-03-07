package main

import (
	"fmt"
)

//定义一个函数类型
type FruitName func() string

type Fruit struct {
	GetFruitName FruitName
	//相当于：
	//GetFruitName func() string
}

func (this *Fruit) DisplayName() {
	fmt.Println(this.GetFruitName())
}
func (this *Fruit) GetName() string {
	return "fruit"
}
func NewFriut() *Fruit {
	f := new(Fruit)
	f.GetFruitName = f.GetName
	return f
}

type Apple struct {
	Fruit
}

func (this *Apple) GetName() string {
	return "apple"
}
func NewApple() *Apple {
	a := new(Apple)
	a.GetFruitName = a.GetName
	return a
}
func main() {
	fruit := NewFriut()
	fruit.DisplayName()
	apple := NewApple()
	apple.DisplayName()
}
