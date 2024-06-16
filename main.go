package main

import (
	"fmt"
)

func binaryInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		fmt.Printf("left =%v, right =%v\n", left, right)

		median := (right + left) / 2

		if nums[median] == target {
			return median
		}

		if nums[median] < target {
			left = median + 1
		} else if nums[median] > target {
			right = median - 1
		}
	}
	return 0
}

func main() {


}
