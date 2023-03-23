package main

import (
	"fmt"
)
type Person struct{
	id int
	name string
	age int
}
func test1(){
	// 使用new创建结构体指针
	var tom=new(Person)
	tom.id=11
	tom.name="dd"
	fmt.Printf("tom: %p\n", tom)
	fmt.Printf("tom: %v\n", *tom)
}

func main() {
	

	tom :=Person{
		id:11,
		name:"dd",
		age:11,
	}

	var p_person *Person
	p_person =&tom
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("p_person: %p\n", p_person)
	fmt.Printf("*p_person: %v\n", *p_person)

	test1()

}