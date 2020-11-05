package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}
