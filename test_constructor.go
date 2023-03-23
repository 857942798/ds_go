package main

import "fmt"

type Person struct{
	name string
}

func NewPsersion(_name string)(*Person,error){
	if _name==""{
		return nil,fmt.Errorf("name is not null")
	}
	return &Person{name:_name},nil
}
func main() {
	per,error:=NewPsersion("ds")
	fmt.Printf("per: %v\n", *per)
	fmt.Printf("error: %v\n", error)
}