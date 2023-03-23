package main

import (
	"fmt"
)

var c=make(chan int)

func main() {
	go func()  {
		for i := 0; i <9; i++ {
			c <- i
		}
		close(c)
	}()

	r:= <- c
	fmt.Printf("r: %v\n", r)
	r = <- c
	fmt.Printf("r: %v\n", r)

	for v:= range c{
		fmt.Printf("v: %v\n", v)
	}

	for{
		v,ok:=<-c
		if ok {
			fmt.Printf("v: %v\n", v)
		}else {
			break
		}
	}

}