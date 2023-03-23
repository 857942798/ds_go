package main

import (
	"fmt"
)

func main(){

	a:=100

	if a>200 {
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}

	switch a{
		case 100: fmt.Println("11")
		default: fmt.Println("22")
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	x:=[...]int {1,2,3}
	for _, v := range x {
		fmt.Printf("x: %v\n",v)
	}

}