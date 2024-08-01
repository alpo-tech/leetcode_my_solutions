package main

import "fmt"

func main() {
	var number uint32 = 4251705699 // пример 32-битного числа (в 16-ричном формате для наглядности)

	// Перевернуть все биты
	invertedNumber := ^number

	// Печать исходного и инвертированного числа в бинарном формате
	fmt.Printf("Original number: %032b\n", number)
	fmt.Printf("Inverted number: %032b\n", invertedNumber)
}
