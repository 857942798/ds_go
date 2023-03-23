package main

import "fmt"


func f1(){
	var num=[...]int{1,2,3}
	for i, v := range num {
		fmt.Printf("%v:%v ", i,v)
	}
}

func f2(){
	var s=[]int{1,2,3}
	for _, v := range s {
		fmt.Printf("v: %v\n", v)
	}
}

func f3(){
	m := make(map[string]string, 0)
	m["name"]="tom"
	m["age"]="20"
	for key,value := range m {
		fmt.Println(key)
		fmt.Println(value)
	}
}

func main(){
	// f1()
	f2()
	f3()
}