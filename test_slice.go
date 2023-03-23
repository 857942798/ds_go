package main

import "fmt"

func test1(){
	var s1 []int
	var s2 []string
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

func test2(){
	var s1=make([]int,2)
	fmt.Printf("s1: %v\n", s1)
}

func test3(){
	var s1=[]int{1,2,4}
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))

	fmt.Printf("s1[0]: %v\n", s1[0])
}

func test4(){
	var s1=[]int{1,2,4}

	s2:=s1[0:1]
	fmt.Printf("s2: %v\n", s2)

}

func test5(){
	var s1=[]int{1,2,3,4,5}
	l :=len(s1)
	for i:=0;i<l;i++{
		fmt.Printf("s1[i]: %v\n", s1[i])
	}
}

func test6(){
	var s1=[]int{1,2,3,4,5}
	for i, v := range s1 {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}

func curd(){
	c1 :=[]int{}
	c1 = append(c1, 1)
	c1 = append(c1, 2)
	c1 = append(c1, 3)
	fmt.Printf("c1: %v\n", c1)

	
	c1 = append(c1[:1], c1[2:]...)
	fmt.Printf("c1: %v\n", c1)

	c1[0]=100
	fmt.Printf("c1: %v\n", c1)
}

func test7(){
	var s1=[]int{1,2,3,4,5}
	s2:=s1
	s2[0]=100
	fmt.Printf("s1: %v\n", s1)

	s3 :=make([]int,3)
	copy(s3,s1)
	s3[1]=200

	fmt.Printf("s1: %v\n", s1)
}

func main(){
	test1()
	test2()
	test3()
	test4()
	test5()
	test6()
	curd()
	test7()
}