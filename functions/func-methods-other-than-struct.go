package main

import "fmt"

func main() {
	var i intIncrementer = 6
	i.incremnt()
	fmt.Println(i)
}

type intIncrementer int

func (i *intIncrementer) incremnt() {
	*i = intIncrementer(int(*i) + 1)
}
