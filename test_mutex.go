package main

import (
	"fmt"
	"sync"
)

var i int

var wt sync.WaitGroup

var lock sync.Mutex

func add() {
	lock.Lock()
	defer wt.Done()
	i+=1
	lock.Unlock()
}

func sub(){
	lock.Lock()
	defer wt.Done()
	i-=1
	lock.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		wt.Add(1)
		go add()
		wt.Add(1)
		go sub()
	}
	wt.Wait()
	fmt.Println("end i:",i)
}