package main

import (
	"fmt"
)

func main() {
	//arrayLine := make([][]int, 0)
	//arrayLine = append(arrayLine, []int{1, 2, 3, 4, 5, 6})
	//fmt.Println(arrayLine)


	test := "testTest   124 % fr"

	for _, val := range test {
		if val >= 65 && val <=90 {
			fmt.Printf("%c", val)
		}
		
	}
}
