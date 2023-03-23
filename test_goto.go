package main

import "fmt"


func f1(){
	i:=4

	if i>=2{
		fmt.Println(i)
	}else{
		goto END
	}
	END: 
		fmt.Println("22")
}


func main(){
	f1()
}

