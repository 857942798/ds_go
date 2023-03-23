package main

import (
	"fmt"
	"time"
)

func showMsg(msg string){
	for i:=0;i<5;i++{
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	go showMsg("java") // go 启动一个协程
	defer showMsg("golang")
	fmt.Println("main end...")// 3 主函数推出,程序就结束了
	
}