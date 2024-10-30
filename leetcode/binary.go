package leetcode

import "fmt"

func reverseBitString(bytes []byte) string {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}

	return string(bytes)
}

func invertBitString(str string) []byte {
	result := make([]byte, len(str))

	for i, bit := range str {
		if bit == '1' {
			result[i] = '0'
		} else {
			result[i] = '1'
		}
	}

	return result
}

func getBitString(x int) string {
	if x == 1 {
		return "0"
	}

	s := getBitString(x - 1)
	return s + "1" + reverseBitString(invertBitString(s))
}

func findKthBit(n int, k int) byte {
	return getBitString(n)[k - 1]
}


func countBits_338(n int) []int {
	result := make([]int, n + 1)

	summaBits := func(value int) int {
		valueStr := fmt.Sprintf("%b", value)
		summa := 0
		for _, num := range valueStr {
			if num == '1' {
				summa += 1
			}
		}
		return summa
	}

	for i := 0; i <= n; i++ {
		result[i] = summaBits(i)
	}

	return result
}
