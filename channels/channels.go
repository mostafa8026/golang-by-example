package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go func() {
		fmt.Println("start receiving data from channel")
		i := <-ch
		fmt.Printf("data has been received, the data is: %v\n", i)
		i = <-ch
		fmt.Printf("data has been received, the data is: %v\n", i)
		wg.Done()
	}()
	go func() {
		fmt.Println("start sending data to the channel")
		ch <- 47
		ch <- 78
		fmt.Println("data has been sent to the channel safely")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("End of the main")
}
