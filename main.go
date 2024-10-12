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




var Romans = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L':50, 'C':100, 'D':500, 'M': 1000}

func romanToInt(s string) int {
	summa := Romans[s[len(s) - 1]]
	for  i := len(s) - 2; i >= 0; i-- {
		if Romans[s[i + 1]] > Romans[s[i]] {
			summa -= Romans[s[i]]
		} else {
			summa += Romans[s[i]]
		}
	}
	return summa
}



func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
