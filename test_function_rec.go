package main

import "fmt"

func f(i int) int{
	if i==1 {
		return 1
	}else{
		return i*f(i-1)
	}
}

func main() {
	fmt.Printf("f(6): %v\n", f(6))
}