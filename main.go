package main

import (
	"fmt"
	"sort"
)

func test(test []int) {
	sort.Ints(test)
	fmt.Println(test)

}

func sortArrayMy(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	min, max := nums[0], nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}
	countElement := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		countElement[nums[i]-min]++
	}

	indexNums := 0
	for i := 0; i < len(countElement); i++ {
		for countElement[i] > 0 {
			nums[indexNums] = i + min
			countElement[i]--
			indexNums++
		}
	}

	return nums

}

func main() {
	//arrayLine := make([][]int, 0)
	//arrayLine = append(arrayLine, []int{1, 2, 3, 4, 5, 6})
	//fmt.Println(arrayLine)

	// test := "testTest   124 % fr"

	// for _, val := range test {
	// 	if val >= 65 && val <=90 {
	// 		fmt.Printf("%c", val)
	// 	}
	//
	// }

	arr1 := []int{1, 2, 0, 4, 1, 2, 3}
	sortar := sortArrayMy(arr1)
	fmt.Println(sortar)
}
