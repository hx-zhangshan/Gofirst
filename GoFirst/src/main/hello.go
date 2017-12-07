package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	fmt.Println("hello world!!")
	test()
	test2()
}
func test() {

	println(a, b, c)
}
func test2() {
	const (
		//		a1 = iota
		//		b1 = "nihaoa"
		//		c1
		//		d1 = iota
		a1 = 1 << iota //值1
		b1 = 3 << iota //值6
		c1             //预计3<<2
		d1             //预计3<<3
	)
	println(a1, b1, c1, d1)
}
