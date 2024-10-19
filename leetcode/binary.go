package leetcode

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
