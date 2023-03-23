package main

import (
	"fmt"
	"os"
)

func createFile(){
	f,err:=os.Create("a.txt")
	if err !=nil {
		fmt.Printf("err: %v\n", err)
	}else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func makeDir(){
	err := os.Mkdir("test", os.ModePerm)
	if err != nil{
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.MkdirAll("a/b/c", os.ModePerm)

	if err2 != nil{
		fmt.Printf("err2: %v\n", err2)
	}
}

func remove(){
	err := os.RemoveAll("a")
	if err !=nil {
		fmt.Printf("err: %v\n", err)
	}
}

func wd(){
	dir, _ := os.Getwd()
	fmt.Printf("dir: %v\n", dir)
	// 切换到指定目录,如果有创建文件的操作,将在该目录操作
	os.Chdir("/Users/dongsheng/home/projects/GolandProjects")
	dir2, _ := os.Getwd()
	fmt.Printf("dir2: %v\n", dir2)
}

func readFile(){
	b2, _ := os.ReadFile("a.txt")
	fmt.Printf("b2: %v\n", string(b2[:]))
}

func writeFile(){
	os.WriteFile("a.txt",[]byte("hello"),os.ModePerm)
}

func main() {
	createFile()
	// makeDir()
	// remove()
	// wd()
	writeFile()
	readFile()
}