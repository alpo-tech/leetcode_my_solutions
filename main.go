package main

import (
	"fmt"
	"leetcode.com/leetcode"
)

func main() {
	//test := []int{-1, -100, 3, 99}
	test := []int{1, 2, 3, 4, 5, 6}

	leetcode.Rotate_189(test, 3)
	fmt.Println(test)
}
