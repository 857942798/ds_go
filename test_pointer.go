package main

import "fmt"


func test1(){
	var ip *int
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %T\n", ip)

	var i int=100
	ip=&i
	fmt.Printf("ip: %v\n", ip)
	fmt.Printf("ip: %v\n", *ip)
}

func test2(){
	a:=[]int{1,2,3}
	var po [3]*int
	for i:=0;i<3;i++{
		po[i]=&a[i]
	}

	for i:=0;i<3;i++{
		fmt.Printf("po[i]: %v\n", *po[i])
	}
}

func main() {
	test1()
	test2()
}