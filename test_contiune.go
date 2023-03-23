package main

import "fmt"

func main(){
	
	for i := 0; i < 10; i++ {
		mylable:
		for j := 0; j < 10; j++ {
			if i==2 && j==2 {
				continue mylable
			}
			fmt.Println("%v,%v",i,j)
		}
	}
}