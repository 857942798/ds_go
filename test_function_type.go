package main

import (
	"fmt"
)

type f func(int,int)int

func sum(a int,b int)(ret int){
	return a+b
}

func sayHello(msg string){
	fmt.Printf("msg: %v\n", msg)
}

func test1(name string,f func(string)){
	f(name)
}


func add(a int,b int) int{
	return a+b
}

func del(a int,b int) int{
	return a-b
}

func get(op string) func(int,int)int{
	switch op{
	case "+": return add
	case "-": return del
	default: return nil
	}
}

func main() {
	var ff f
	ff=sum
	fmt.Printf("ff(1, 2): %v\n", ff(1, 2))

	test1("ds",sayHello)

	oo:=get("+")
	fmt.Printf("oo(1, 2): %v\n", oo(1, 2))

	oo=get("-")
	fmt.Printf("oo(1, 2): %v\n", oo(1, 2))


	max :=func (a int,b int)int  {
		if(a>b){
			return a
		}else{
			return b
		}
	}
	d:=max(2,3)
	fmt.Printf("d: %v\n", d)


	mm :=func (a int,b int)int  {
		if(a>b){
			return a
		}else{
			return b
		}
	}(2,4)
	fmt.Printf("mm: %v\n", mm)
	
}