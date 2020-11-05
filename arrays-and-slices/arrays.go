package main

import (
	"fmt"

	"github.com/bradhe/stopwatch"
)

func main() {
	a := [2]int{1, 2}
	var b [2]int = [2]int{1, 2}
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(c))

	f := c
	f[2] = 89
	fmt.Println(c)
	fmt.Println(f)

	sl := []int{1, 2, 3, 4}
	fmt.Println(sl)
	sl = append(sl, 5)
	fmt.Println(sl)

	sll := sl
	sll[2] = 89
	fmt.Println(sl)
	fmt.Println(sll)

	news := []int{}
	fmt.Printf("%v len %v cap %v\n", news, len(news), cap(news))
	news = append(news, 1)
	fmt.Printf("%v len %v cap %v\n", news, len(news), cap(news))
	news = append(news, 2)
	fmt.Printf("%v len %v cap %v\n", news, len(news), cap(news))
	news = append(news, 3)
	fmt.Printf("%v len %v cap %v\n", news, len(news), cap(news))
	news = append(news, 4, 5)
	fmt.Printf("%v len %v cap %v\n", news, len(news), cap(news))

	watch := stopwatch.Start()
	ms := make([]int, 0, 1)
	for i := 0; i < 100000000; i++ {
		ms = append(ms, i)
	}
	watch.Stop()
	fmt.Printf("time: %v\n", watch.Milliseconds())
	fmt.Printf("len %v cap %v\n", len(ms), cap(ms))

	watch.Stop()
	watch.Start()
	ms2 := make([]int, 0, 100000000)
	for i := 0; i < 100000000; i++ {
		ms2 = append(ms2, i)
	}
	watch.Stop()
	fmt.Printf("time: %v\n", watch.Milliseconds())
	fmt.Printf("len %v cap %v\n", len(ms2), cap(ms2))

}
