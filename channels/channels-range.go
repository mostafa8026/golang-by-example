package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func() {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}()
	go func() {
		defer func() {
			close(ch)
		}()
		ch <- 42
		ch <- 27
		wg.Done()
	}()
	wg.Wait()
}
