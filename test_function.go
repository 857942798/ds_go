package main

import "fmt"

func sum(a int,b int)(ret int){
	return a+b
}

func test1()(name string,age int){
	// return "dtd",1
	name ="2"
	age=1
	return
}

func test2(a int){
	a=200
}

func test3(s []int){
	s[0]=100
}

func test4(ages ...int){
	for i := 0; i < len(ages); i++ {
		fmt.Printf("ages[i]: %v\n", ages[i])
	}

}


func main() {
	r:=sum(1,2)
	fmt.Printf("r: %v\n", r)
	fmt.Println(test1())

	a:=100
	test2(a)
	fmt.Printf("a: %v\n", a)

	s:=[]int{1,2,3}
	test3(s)
	fmt.Printf("s[0]: %v\n", s[0])

	test4(1,2,3,4)
}