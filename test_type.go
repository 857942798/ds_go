package main

import "fmt"

func test1(){
	// 类型别名 
	type myint=int
	var i myint
	i=100
	fmt.Printf("i: %T\n", i) // int

	// 类型定义
	type myint2 int
	var j myint2
	j=100
	fmt.Printf("j: %T\n", j) //main.myint2
}

func main() {
	test1()
}