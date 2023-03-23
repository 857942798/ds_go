package main

import (
	"fmt"
	"sync/atomic"
)

func add(){
	var i int32=100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i: %v\n", i)

	var j int64=200
	atomic.AddInt64(&j,1)
	fmt.Printf("j: %v\n", j)

}

func load(){
	var i int32=100
	var j=atomic.LoadInt32(&i) //read
	fmt.Printf("j: %v\n", j) 

	atomic.StoreInt32(&i,2022)
	fmt.Printf("i: %v\n", i) //write

}

func cas(){
	var i int32=100
	atomic.CompareAndSwapInt32(&i,100,200)
	fmt.Printf("i: %v\n", i)
}

func main() {
	add()
	load()
	cas()
}