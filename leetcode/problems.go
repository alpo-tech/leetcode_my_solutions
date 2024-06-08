package leetcode

import (
	"strconv"
)

func TwoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for index, value := range nums {
		i, ok := result[target-value]
		if ok != false {
			return []int{i, index}
		}
		result[value] = index
	}
	return []int{0, 0}
}

func IsPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	for index := 0; index < len(strX)/2; index++ {
		if strX[index] != strX[len(strX)-index-1] {
			return false
		}
	}
	return true
}

func IsPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}

	reverseX := 0
	copyX := x

	// сравниваем первую цифру и последнюю
	for copyX > 0 {
		reverseX = reverseX*10 + copyX%10
		copyX = copyX / 10
	}

	return reverseX == x
}
func RomanToInt(s string) int {
	b := map[int32]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	previosValue := 10000000

	for _, value := range s {
		currenValue := b[value]
		if currenValue > previosValue {
			result -= previosValue
			result += (currenValue - previosValue)
		} else {
			result += currenValue
		}

		previosValue = currenValue
	}

	return result
}
