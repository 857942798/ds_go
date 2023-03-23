package main

import (
	"fmt"
	"io"
	"os"
)

// 打开关闭文件
func OpenClose(){
	f5, err := os.Open("a.txt")
	if err !=nil {
		fmt.Printf("err: %v\n", err)
	}else{
		fmt.Printf("f5.Name(): %v\n", f5.Name())
		f5.Close()
	}
	// 不存在就从创建一个新的
	os.OpenFile("a1.txt",os.O_RDONLY|os.O_CREATE,755)

}

// 创建文件
func  createFile(){
	os.Create("a2.txt")
	os.CreateTemp("","tmp")
}

// read
func readOps(){
	// 循环读取
	f5, _ := os.Open("a.txt")
	for{
		// 开启一个缓冲区,每次读三个字节
		buf :=make([]byte,3)
		_,err:=f5.Read(buf)
		if err ==io.EOF{
			break
		}
		fmt.Printf("buf: %v\n", string(buf))
	}
	
	f5.Close()

	f6, _ := os.Open("a.txt")
	buf2 :=make([]byte,3)
	// 从某个偏移量开始读取
	f6.ReadAt(buf2, 3)
	fmt.Printf("buf2: %v\n", string(buf2))

	f6.Close()

	de, _ := os.ReadDir("user")
	// 遍历目录
	for _, v := range de {
		fmt.Printf("v.IsDir(): %v\n", v.IsDir())
		fmt.Printf("v.Name(): %v\n", v.Name())
	}
}

func main() {
	// OpenClose()
	// createFile()
	readOps()
}