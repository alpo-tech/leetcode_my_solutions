package main

import "fmt"

func main() {
	var a, b byte
	a = 255
	b = 48

	fmt.Printf("%08b = %d\n", a, a)
	c := a & b
	fmt.Println(c)

	a = a << 7
	//a = a | 1
	fmt.Println(a)
}
