package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsIndex := make(map[int]int)
	for index, num := range nums {
		if secondIndex, ok := numsIndex[target-num]; ok {
			return []int{index, secondIndex}
		}
		numsIndex[num] = index
	}

	return []int{}
}

func main() {

	tests := []struct {
		nums   []int
		target int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9},
		{nums: []int{3, 2, 4}, target: 6},
		{nums: []int{3, 3}, target: 6},
	}

	for _, test := range tests {
		fmt.Println(twoSum(test.nums, test.target))
	}
}
