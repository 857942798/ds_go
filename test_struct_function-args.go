package main

import (
	"fmt"
)

type Person struct{
	id int
	name string
}

func showPerson(per Person){
	per.id=1
	per.name="kk"
	fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person){
	per.id=2
	per.name="kk2"
	fmt.Printf("per: %v\n", per)
}

func main() {
	tom:=Person{
		id:100,
		name:"dd",
	}
	fmt.Printf("tom: %v\n", tom)
	// 这里是值传递,相当于拷贝了一个副本,方法里对对象的修改不会影响到外面
	showPerson(tom)
	fmt.Printf("tom: %v\n", tom)

	per:=&tom
	// 传递一个指针进去,则是引用传递
	showPerson2(per)
	fmt.Printf("tom: %v\n", tom)
}