package leetcode

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type stackInt struct {
	stack []int
	lock  sync.RWMutex
}

func (stack *stackInt) Push(value int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.stack = append(stack.stack, value)
}

func (stack *stackInt) Pop() (int, error) {
	if stack.Size() > 0 {
		stack.lock.Lock()
		defer stack.lock.Unlock()
		value := stack.stack[stack.Size()-1]
		stack.stack = stack.stack[:stack.Size()-1]
		return value, nil
	}
	return 0, fmt.Errorf("Size stack = 0")
}

func (stack *stackInt) Empty() bool {
	return stack.Size() == 0
}

func (stack *stackInt) Front() (int, error) {
	if stack.Size() > 0 {
		stack.lock.Lock()
		defer stack.lock.Unlock()
		value := stack.stack[stack.Size()-1]
		return value, nil
	}
	return 0, fmt.Errorf("Size stack = 0")
}

func (stack *stackInt) Size() int {
	return len(stack.stack)
}

func newStack() stackInt {
	return stackInt{
		stack: make([]int, 0),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(list *ListNode) {
	for l := list; l != nil; l = l.Next {
		fmt.Printf("%v -> ", l.Val)
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var current *ListNode
	var result *ListNode = &ListNode{}
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

func TwoSum(nums []int, target int) []int {
	result := make(map[int]int)
	for index, value := range nums {
		i, ok := result[target-value]
		if ok != false {
			return []int{i, index}
		}
		result[value] = index
	}
	return []int{0, 0}
}

func IsPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	for index := 0; index < len(strX)/2; index++ {
		if strX[index] != strX[len(strX)-index-1] {
			return false
		}
	}
	return true
}

func IsPalindromeV2(x int) bool {
	if x < 0 {
		return false
	}

	reverseX := 0
	copyX := x

	// сравниваем первую цифру и последнюю
	for copyX > 0 {
		reverseX = reverseX*10 + copyX%10
		copyX = copyX / 10
	}

	return reverseX == x
}
func RomanToInt(s string) int {
	b := map[int32]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	previosValue := 10000000

	for _, value := range s {
		currenValue := b[value]
		if currenValue > previosValue {
			result -= previosValue
			result += (currenValue - previosValue)
		} else {
			result += currenValue
		}

		previosValue = currenValue
	}

	return result
}

func LongestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]
	for _, str := range strs[1:] {
		sizePrefix := 0
		if len(str) > len(commonPrefix) {
			sizePrefix = len(commonPrefix)
		} else {
			sizePrefix = len(str)
			commonPrefix = commonPrefix[:sizePrefix]
		}

		for index, _ := range str[:sizePrefix] {
			if commonPrefix[index] != str[index] {
				commonPrefix = str[:index]
				break
			}
		}

	}

	return commonPrefix
}

func ValidParentheses(str string) bool {

	reverse := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := make([]int32, 0, 10)

	for _, skobka := range str {
		if skobka == '(' || skobka == '[' || skobka == '{' {
			stack = append(stack, skobka)
		} else if skobka == ')' || skobka == ']' || skobka == '}' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if reverse[last] != skobka {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func ValidParenthesesV2(str string) bool {

	customStack := newStack()

	for _, value := range str {
		switch value {
		case '{':
			customStack.Push(int(value))
			break
		case '[':
			customStack.Push(int(value))
			break
		case '(':
			customStack.Push(int(value))
			break
		case '}':
			valueStack, err := customStack.Pop()
			if err != nil || valueStack != '{' {
				return false
			}
			break
		case ']':
			valueStack, err := customStack.Pop()
			if err != nil || valueStack != '[' {
				return false
			}
			break
		case ')':
			valueStack, err := customStack.Pop()
			if err != nil || valueStack != '(' {
				return false
			}
			break
		default:
			return false
		}
	}

	return customStack.Empty()
}

func RemoveDuplicatess(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[k] {
			k++
			nums[k] = nums[i]
		} else {
			nums[k] = nums[i]
		}
	}

	return k + 1
}

func RemoveElement(nums []int, val int) int {
	len := len(nums)
	count := 0

	for i := 0; i < len; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	return len - (len - count)
}

func strStrSample(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func StrStr(haystack string, needle string) int {
	// O(N)
	prefix := make([]int, len(needle))
	j := 0
	for i := 1; i < len(needle); {
		if needle[i] == needle[j] {
			prefix[i] = j + 1
			i++
			j++
		} else {
			if j == 0 {
				prefix[i] = 0
				i++
			} else {
				j = prefix[j-1]
			}
		}
	}

	k := 0
	l := 0
	for k < len(haystack) {
		if haystack[k] == needle[l] {
			k++
			l++
			if l == len(needle) {
				return k - l
			}
		} else {
			if l == 0 {
				k++
			} else {
				l = prefix[l-1]
			}
		}
	}

	return -1
}
