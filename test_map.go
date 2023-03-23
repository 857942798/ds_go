package main

import (
	"fmt"
)

func test2(){
	m2:=make(map[string]string)
	m2["name"]="12ddfdf"
	m2["age"]="22"

	var key="name"
	var value=m2[key]
	fmt.Printf("value: %v\n", value)

	v,ok :=m2["name"]
	fmt.Printf("ok: %v\n", ok)
	fmt.Printf("v: %v\n", v)
}

func test1(){
	var m map[string]string
	m=make(map[string]string)
	fmt.Printf("m: %v\n", m)


	var m1=map[string]string{"name":"tom","age":"11"}
	fmt.Printf("m1: %v\n", m1)


	m2:=make(map[string]string)
	m2["name"]="hahah"
	m2["age"]="22"
	fmt.Printf("m2: %v\n", m2)

}

func main(){
	test1()
	test2()
}