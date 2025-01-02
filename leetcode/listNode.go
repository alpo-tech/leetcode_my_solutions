package leetcode

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(list *ListNode) {
	for l := list; l != nil; l = l.Next {
		fmt.Printf("%v ", l.Val)
	}
}

func CreateList(nums []int) *ListNode {
	head := &ListNode{}
	head.Val = nums[0]
	res := head
	for _, value := range nums[1:] {
		head.Next = &ListNode{}
		head = head.Next
		head.Val = value
	}
	return res
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

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	listNodesMap := make(map[*ListNode]int, 0)
	for i := headA; i != nil; i = i.Next {
		listNodesMap[i]++
	}

	for i := headB; i != nil; i = i.Next {
		value := listNodesMap[i]
		if value != 0 {
			return i
		}
	}

	return nil
}

func ReverseList(head *ListNode) *ListNode {
	arrayListNode := make([]*ListNode, 0)

	current := head

	for current != nil {
		arrayListNode = append(arrayListNode, current)
		current = current.Next
	}

	for i := len(arrayListNode) - 1; i > 0; i-- {
		arrayListNode[i].Next = arrayListNode[i-1]
	}

	if len(arrayListNode) == 0 {
		return nil
	}

	arrayListNode[0].Next = nil
	return arrayListNode[len(arrayListNode)-1]
}

func ReverseList_206(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	currenList := head.Next
	previosList := head
	previosList.Next = nil
	for currenList != nil {
		nextList := currenList.Next
		currenList.Next = previosList
		previosList = currenList
		currenList = nextList
	}
	return previosList
}

func reverseListCount_92(head *ListNode, count int) (*ListNode, *ListNode) {
	var revHead *ListNode
	origHead := head

	for count > 0 {
		tmp := head.Next
		head.Next = revHead
		revHead = head
		head = tmp
		count--
	}

	return origHead, revHead
}

func ReverseBetween_92(head *ListNode, left int, right int) *ListNode {

	var beforeLeft, newEnd, newBegin, afterRight *ListNode

	result := head

	for i := 1; i <= right; i++ {
		if i == left {
			leftRev := head
			for i <= right {
				afterRight = head.Next
				head = head.Next
				i++
			}
			newEnd, newBegin = reverseListCount_92(leftRev, right-left+1)
			break
		}
		beforeLeft = head
		head = head.Next
	}

	if beforeLeft == nil {
		result = newBegin
	} else {
		beforeLeft.Next = newBegin
	}
	newEnd.Next = afterRight
	return result
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{0, head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	current := prev.Next

	for i := 0; i < right-left; i++ {
		nextNode := current.Next
		current.Next = nextNode.Next
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}

func getDecimalValue_1290(head *ListNode) int {
	arraybit := make([]int, 0)

	for head != nil {
		arraybit = append(arraybit, head.Val)
		head = head.Next
	}

	result := 0
	for i := len(arraybit) - 1; i >= 0; i-- {
		dig := len(arraybit) - i - 1
		result += int(arraybit[i] * int(math.Pow(2, float64(dig))))
	}

	return result
}

func middleNode_876(head *ListNode) *ListNode {
	first := head
	second := head

	for second != nil {
		second = second.Next
		if second == nil {
			break
		}
		first = first.Next
		second = second.Next
	}
	return first
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r1 := make([]int, 0)
	r2 := make([]int, 0)

	i1 := l1
	for i1 != nil {
		r1 = append(r1, i1.Val)
		i1 = i1.Next
	}

	i2 := l2
	for i2 != nil {
		r2 = append(r2, i2.Val)
		i2 = i2.Next
	}

	if len(r1) > len(r2) {
		for len(r1) != len(r2) {
			r2 = append(r2, 0)
		}
	}

	if len(r1) < len(r2) {
		for len(r1) != len(r2) {
			r1 = append(r1, 0)
		}
	}

	r := make([]int, len(r1))
	value := 0
	for i := 0; i < len(r1); i++ {
		value += r1[i] + r2[i]
		r[i] = value % 10
		value /= 10
	}

	if value != 0 {
		r = append(r, value)
	}

	return CreateList(r)

}

func RemoveNthFromEnd_19(head *ListNode, n int) *ListNode {
	countNode := 0
	begin := head
	for begin != nil {
		countNode += 1
		begin = begin.Next
	}

	removeNth := countNode - n
	if removeNth == 0 {
		return head.Next
	}
	begin = head
	previous := head
	for removeNth != 0 {
		removeNth--
		previous = begin
		begin = begin.Next
	}

	previous.Next = begin.Next
	return head
}

func SwapPairs_24(head *ListNode) *ListNode {

	first := head
	if first == nil {
		return head
	}

	second := head.Next
	if second == nil {
		return head
	}

	prev := &ListNode{-1, head}
	result := prev
	for first != nil && second != nil {
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		first = first.Next
		if first != nil {
			second = first.Next
		}
	}

	return result.Next
}

func rotateHelper(head *ListNode) *ListNode {
	begin := head
	previous := &ListNode{-1, begin}
	for begin.Next != nil {
		previous = begin
		begin = begin.Next
	}

	previous.Next = nil
	begin.Next = head
	return begin
}

func RotateRight(head *ListNode, k int) *ListNode {
	newHead := head
	if newHead == nil {
		return newHead
	}

	n := 1
	for head.Next != nil {
		head = head.Next
		n++
	}

	k = k % n
	for k > 0 {
		k--
		newHead = rotateHelper(newHead)
	}

	return newHead
}

func DeleteDuplicatesFromList_83(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head
	second := head.Next

	for second.Next != nil {
		if first.Val == second.Val {
			second = second.Next
		} else {
			first.Next = second
			first = second
			second = second.Next
		}
	}

	if first.Val == second.Val {
		first.Next = nil
	} else {
		first.Next = second
	}

	return head
}

func DeleteDuplicatesFromList_82(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	previous := &ListNode{-1, nil}
	result := previous
	first := head
	second := head.Next
	checkLast := true

	for second != nil {
		if first.Val != second.Val {
			previous.Next = first
			previous = first
			previous.Next = nil
			first = second
			second = second.Next
		} else {
			previous.Next = nil
			for second != nil && first.Val == second.Val {
				second = second.Next
			}
			if second != nil {
				checkLast = true
				first = second
				second = second.Next
			} else {
				checkLast = false
			}
		}
	}

	if previous.Val != first.Val && checkLast {
		previous.Next = first
		if first.Next != nil && first.Val == first.Next.Val {
			first.Next = nil
		}
	}

	return result.Next
}

func Partition_86(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	before := &ListNode{}
	after := &ListNode{}

	curr_before := before
	curr_after := after

	for head != nil {
		if head.Val < x {
			curr_before.Next = head
			curr_before = curr_before.Next
		} else {
			curr_after.Next = head
			curr_after = curr_after.Next
		}
		head = head.Next
	}

	curr_before.Next = after.Next
	curr_after.Next = nil

	return before.Next
}
