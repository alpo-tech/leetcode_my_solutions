package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	testSlice(slice)
	fmt.Println(slice)
}

func testSlice(slice []int) {
	slice[0] = 5
	slice[4] = 1
}
