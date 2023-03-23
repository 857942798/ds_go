package main

import "fmt"

func main() {
	type Dog struct{
		name string
		age int
	}
	type Person struct{
		dog Dog
		name string
		id int	
	}
	dog:=Dog{
		name:"huahua",
		age:2,
	}
	per:=Person{
		dog:dog,
		name: "tom",
		id: 10,
	}
	fmt.Printf("per: %v\n", per)

	fmt.Printf("per.dog: %v\n", per.dog)
}