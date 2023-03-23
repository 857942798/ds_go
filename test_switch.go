package main

import "fmt"

func main(){
	grade:="A"
	switch grade {
		case "A": fmt.Println("A")
		case "B": fmt.Println("B")
		default: fmt.Println("C")
	}

	day:=1
	switch day {
		case 1,2,3,4,5: 
			fmt.Println("工作日")
			fallthrough
		case 6,7: 
			fmt.Println("休息日")
		default : 
			fmt.Println("非法输入")
		
	}
}