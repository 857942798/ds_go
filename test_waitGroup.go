package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMsg(i int){
	defer wp.Done()
	fmt.Printf("i: %v\n", i)
}

func main() {
	for i := 0; i < 10; i++ {
		// 启动一个协程来执行
		wp.Add(1)
		go showMsg(i)
	}
	// 计数器减到0才继续向下执行
	wp.Wait()
	// 主协程
	fmt.Println("end....")
}