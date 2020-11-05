package main

import "fmt"

func main() {
	s := "<testing the strings>"

	var ss string = "سیبشسیب"
	fmt.Printf("%v %T\n", s, s)
	fmt.Printf("%v %T\n", ss[5], ss[5])

	bb := []byte(ss)
	fmt.Println(bb)

	r := 't'
	fmt.Printf("%v %T\n", r, r)
}
