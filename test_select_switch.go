package main

import (
	"fmt"
	"time"
)

var chanInt=make(chan int)
var chanStr=make(chan string)

func main() {
	go func(){
		chanInt <- 100
		chanStr <- "hello"
		// defer close(chanInt)
		// defer close(chanInt)
	}()
	for{
		select	{
			case r:= <- chanInt:
				fmt.Printf("r: %v\n", r)
			case r:= <- chanStr:
				fmt.Printf("r: %v\n", r)
				// 如果没写default会发生死锁
			default:
				fmt.Println("default...")
		}
		time.Sleep(time.Second)
	}
}