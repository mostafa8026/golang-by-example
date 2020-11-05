package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello anonymously"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye anonymously"
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello!")
}
