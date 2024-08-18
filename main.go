package main

import "leetcode.com/leetcode"

func main() {
	list := leetcode.CreateList([]int{1})

	//leetcode.PrintList(list)

	revList := leetcode.ReverseBetween_92(list, 2, 4)

	leetcode.PrintList(revList)
}
