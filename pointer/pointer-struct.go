package main

import "fmt"

func main() {
	var ms *myStruct
	ms = new(myStruct)
	useStruct(ms)
	fmt.Println(ms.foo)
	ms.foo = 88
	fmt.Println(ms.foo)
}

func useStruct(ms *myStruct) {
	ms.foo = 23
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}
