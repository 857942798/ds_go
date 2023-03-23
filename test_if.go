package main

import "fmt"

func main(){
	a:=1
	b:=2
	// gofmt
	if a>b{
		fmt.Println(a)
	}else {
		fmt.Println(b)
	}

	if age:=20;age>18 {
		fmt.Println("1")
	}else{
		fmt.Println("2")
	}

	// var name string
	// fmt.Scan(&name)
	// fmt.Printf("name: %v\n", name)

	// 多重选择 if else if
	score := 80
	if score>=60 && score<=70{
		fmt.Printf("C")
	}else if score<60{
		fmt.Printf("D")
	}else{
		fmt.Printf("A")
	}
}