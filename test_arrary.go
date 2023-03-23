package main

import "fmt"

func main() {
	var a1 [2]int 
	var a2 [3]string
	var a3 [2]bool
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)

	// 数组初始化
	var a4 =[2]int{1,2}
	fmt.Printf("a4: %v\n", a4)

	var a5=[...]string{"sds","22"}
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a5: %v\n", len(a5))

	var a6 =[...]bool{1:true,5:true}
	fmt.Printf("a6: %v\n", a6)

	var a10 [2]int
	a10[0]=1
	a10[1]=2
	fmt.Printf("a10[0]: %v\n", a10[0])
	fmt.Printf("a10[1]: %v\n", a10[1])
	fmt.Printf("len(a10): %v\n", len(a10))

	for i:=0;i<len(a10);i++{
		fmt.Printf("a10[i]: %v\n", a10[i])
	}

	for _, v := range a10 {
		fmt.Printf("v: %v\n", v)
	}
}