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

	for i, j := 0, len(newStr) - 1; i < j; i, j = i+1, j -1 {
		newStr[i], newStr[j] = newStr[j], newStr[i]
	}

	return string(newStr)
}
