package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32=100

// cas compare and swap old new
func add(){
	// 通过原子变量实现先线程安全
	atomic.AddInt32(&i,1)
}

func sub(){
	atomic.AddInt32(&i,-1)
}

func main() {
	for i := 0; i < 1000; i++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second*2)

	fmt.Printf("i: %v\n", i)
}