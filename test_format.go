package main

import "fmt"

type WebSite struct {
	Name string
}

func main() {
	// %v var
	fmt.Printf("\"hello\": %v\n", "hello")
	site := WebSite{Name: "ds"}
	fmt.Printf("site: %v\n", site)
	// %#v 
	fmt.Printf("site: %#v\n", site)
	// %T 
	fmt.Printf("site: %T\n", site)

	b:=true
	fmt.Printf("b: %t\n", b)

	i:=8
	fmt.Printf("i: %v\n", i)
	fmt.Printf("i: %b\n", i)
	i=64
	fmt.Printf("i: %c\n", i)

	x :=100
	p :=&x
	fmt.Printf("i: %p\n", p)
}
