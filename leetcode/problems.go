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

func LongestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]
	for _, str := range strs[1:] {
		sizePrefix := 0
		if len(str) > len(commonPrefix) {
			sizePrefix = len(commonPrefix)
		} else {
			sizePrefix = len(str)
			commonPrefix = commonPrefix[:sizePrefix]
		}

		for index, _ := range str[:sizePrefix] {
			if commonPrefix[index] != str[index] {
				commonPrefix = str[:index]
				break
			}
		}

	}

	return commonPrefix
}

func ValidParentheses(str string) bool {

	reverse := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]int32, 0, 10)

	for _, skobka := range str {
		if skobka == '(' || skobka == '[' || skobka == '{' {
			stack = append(stack, skobka)
		} else if skobka == ')' || skobka == ']' || skobka == '}' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if reverse[last] != skobka {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
