package main

import (
	"fmt"

	"leetcode.com/leetcode"
)

func main() {

	test := []int{1, 1, 1, 2, 2, 3}
	test1 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	leetcode.RemoveDuplicates_80(test)
	leetcode.RemoveDuplicates_80(test1)
	fmt.Println(test)
	fmt.Println(test1)
}
