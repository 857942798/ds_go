package main

import "fmt"

 var i int=valInit()

func init() {
	fmt.Println("init...")
}

func valInit() int{
	fmt.Println("val init...")
	return 100
}

func main() {
	fmt.Println("main....")
}