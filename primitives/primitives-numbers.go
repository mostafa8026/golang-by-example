package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = 127  // -128 to 127
	var b uint8 = 255 // 0 to 255
	var c uint64 = 1844674407370955161
	var d int = 1844674407370955161

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	e := unsafe.Sizeof(d)
	fmt.Println(e)
	var f int8 = 120
	fmt.Printf("remainder of a %% f: %v\n", (a % f))

	var fa float32 = 3.14
	fb := 3.14e19
	fmt.Printf("%v, %T\n", fa, fa)
	fmt.Printf("%v, %T\n", fb, fb)
}
