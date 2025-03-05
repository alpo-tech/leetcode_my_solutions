package leetcode

func removeStars(s string) string {
	countStars := 0
	newStr := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '*' && countStars == 0 {
			newStr = append(newStr, s[i])
		} else if s[i] == '*' {
			countStars++
		} else {
			countStars--
		}
	}

	for i, j := 0, len(newStr)-1; i < j; i, j = i+1, j-1 {
		newStr[i], newStr[j] = newStr[j], newStr[i]
	}

	return string(newStr)
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	stack = append(stack, asteroids[0])
	getMod := func(number int) int {
		if number < 0 {
			return -number
		}
		return number
	}
	for i := 1; i < len(asteroids); i++ {
		if len(stack) != 0 {
			current := stack[len(stack)-1]
			if (current > 0 && asteroids[i] > 0) || (current < 0 && asteroids[i] < 0) {
				stack = append(stack, asteroids[i])
			} else if current > 0 && getMod(asteroids[i]) > getMod(current) {
				i--
				stack = stack[:len(stack)-1]
			} else if current > 0 && getMod(asteroids[i]) == getMod(current) {
				stack = stack[:len(stack) - 1]
			} else if current < 0 {
				stack = append(stack, asteroids[i])
			}
		} else {
			stack = append(stack, asteroids[i])
		}
	}
	return stack
}
