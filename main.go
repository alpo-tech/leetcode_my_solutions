package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	result := make([]byte, len(word1)+len(word1))

	i, j, k := 0, 0, 0
	for ; i < len(word1) && j < len(word1); i, j = i+1, j+1 {
		result[k] = word1[i]
		k++
		result[k] = word2[j]
		k++
	}

	for i < len(word1) {
		result[k] = word1[i]
		i++
		k++
	}

	for j < len(word1) {
		result[k] = word2[j]
		j++
		k++
	}

	return string(result)

}

func main() {
	fmt.Println("test")
}
