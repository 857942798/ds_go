package main

import "fmt"

type Person struct{
	name string
}
// 属性和方法分开来写
// (per Person) 接受者
func (per Person) eat(){
	per.name="ddd"
	fmt.Println("%v,eat....",per.name)
}
func (per Person) sleep(){
	fmt.Println("%v,sleep....",per.name)
}

func (per *Person) sleep2(){
	// 自动解引用
	per.name="hh"
	fmt.Println("%v,sleep....",per.name)
}

func (per Person) login(name string,pwd string)bool{
	if name=="ds" {
		return true
	}
	return false
}

func main() {
	per:=Person{
		name: "ds",
	}
	// struct结合他的方法就等价于面向对象中的类
	per.eat()
	fmt.Printf("per.name: %v\n", per.name)

	per.sleep()
	fmt.Println("per.login(\"ds\", \"123\"): %v\n", per.login("ds", "123"))

	per2:=&Person{
		name: "ds",
	}
	per2.sleep2()
	fmt.Printf("per2.name: %v\n", per2.name)
}
