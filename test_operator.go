package main

import "fmt"

func getA() int {
	return 122
}

func main(){
	a :=100
	b :=100
	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a * b): %v\n", (a * b))

	r:=a==b
	fmt.Printf("r: %v\n", r)

	a=getA()
	fmt.Printf("a: %v\n", a)

	a=4 //100
	fmt.Printf("a: %v\n", a<<1)

}