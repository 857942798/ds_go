package main

import (
		"fmt"
)

func add() func(int)int{
	var x int
	return func(i int) int {
		x+=i
		return x
	}
}

func cal(base int)(func(int)int ,func(int)int){
	add:=func (a int)int  {
		 base+=a
		 return base
	}
	del:= func (b int)int  {
		base-=b
		return base
	}
	return add,del
}


func main() {
	f:=add()
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(20): %v\n", f(20))

	r:=add()
	fmt.Printf("r(10): %v\n", r(10))
	fmt.Printf("r(100): %v\n", r(100))

	aa,dd:=cal(100);
	fmt.Printf("add(100): %v\n", aa(100))
	fmt.Printf("del(50): %v\n", dd(50))
	
}