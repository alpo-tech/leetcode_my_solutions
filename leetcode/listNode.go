package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(list *ListNode) {
	for l := list; l != nil; l = l.Next {
		fmt.Printf("%v -> ", l.Val)
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	var result = &ListNode{}
	current := result
	current.Val = -101
	for head != nil {
		if head.Val != current.Val {
			current.Next = head
			current = current.Next
		}
		head = head.Next
	}

	current.Next = nil

	return result.Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var current *ListNode
	var result = &ListNode{}
	current = result

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return result.Next
}

func HasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	first := head.Next
	if first == nil {
		return false
	}

	second := first.Next
	if second == nil {
		return false
	}

	for first != nil && second != nil {
		if first == second {
			return true
		}

		first = first.Next
		if first == nil {
			return false
		}

		second = second.Next
		if second == nil {
			return false
		}

        second = second.Next 

	}

	return false
}
