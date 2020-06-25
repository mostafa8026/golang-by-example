package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//var msg = "Hello"
	for i := 1; i < 20; i++ {
		wg.Add(2)
		go func1()
		go func2()
	}

	wg.Wait()
}

func func1() {
	fmt.Println("func 1")
	wg.Done()
}

func func2() {
	fmt.Println("func 2")
	wg.Done()
}
