package main

import "fmt"

const (
	_ = 1 << (10 * iota)
	zahedan
	mashad
	ahvaz
	khorasan
)

func main() {
	const a int = 5
	const b = 90.56

	var city int
	fmt.Println(city == zahedan)

	fmt.Println(a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Println(zahedan)
	fmt.Println(mashad)
	fmt.Println(ahvaz)
	fmt.Println(khorasan)
}
