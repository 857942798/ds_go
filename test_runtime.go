package main

import (
	"fmt"
	"runtime"
	"time"
)

func showMsg(msg string){
	for i:=0;i<5;i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func show(msg string){
	for i:=0;i<10;i++ {
		fmt.Printf("j: %v\n", i)
		if i>5{
			runtime.Goexit()
		}
	}
}

func a(){
	for i:=0;i<10;i++ {
		fmt.Printf("a: %v\n", i)
	}
}

func b(){
	for i:=0;i<10;i++ {
		fmt.Printf("b: %v\n", i)
	}
}

func main() {
	go showMsg("java")

	for i:=0;i<2;i++{
		runtime.Gosched() // 让出cpu时间片,让其它子协程执行
		fmt.Println("Gosched...")
	}
	fmt.Println("end...")

	go show("goexit")

	time.Sleep(time.Second)

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}