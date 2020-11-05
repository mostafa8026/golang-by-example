package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		fmt.Println(<-ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
