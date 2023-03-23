package main

import "fmt"

// 自定义一个类型
type Person struct{
	id int
	name string
	age int
}

type Customer struct{
	id,age int
	name,email string
}

// go 没有对象的概念,但是可以使用结构体实现,可以通过结构体嵌套来实现继承的效果
func main() {
	var tom Person
	tom.age=18
	tom.name="ds"
	tom.id=1
	fmt.Printf("tom: %v\n", tom)

	// 匿名结构体
	var my struct{
		id int
		name string
	}
	my.id=1
	my.name="hh"
	fmt.Printf("my: %v\n", my)

	// 结构体的初始化,键值对、列表、部分初始化
	var mm Person
	mm=Person{
		id : 1,
		name: "dd",
	}
	fmt.Printf("mm: %v\n", mm)
}