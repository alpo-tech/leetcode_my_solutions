package main

import "fmt"

func main() {
	number := 27

	ascii := 65

	result := make([]rune, 0)
	for number > 0 {
		ostatok := number % 27
		result = append(result, (rune(ostatok + ascii)))
		number = number / 27
	}

	fmt.Println(string(result))

}
