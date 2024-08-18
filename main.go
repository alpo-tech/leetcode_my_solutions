package main

import "leetcode.com/leetcode"

func main() {
	list := leetcode.CreateList([]int{1, 2, 3, 4, 5})

	//leetcode.PrintList(list)

	revList := leetcode.ReverseBetween(list, 2, 4)

	leetcode.PrintList(revList)
}
