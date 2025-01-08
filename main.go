package main

import (
	"fmt"
)

func reverseBitString(bytes []byte) string {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}

	return string(bytes)
}

type List struct {
	val  int
	next *List
}

type MyHashSet struct {
	set *List
}

func Constructor() MyHashSet {
	return MyHashSet{&List{-1, nil}}
}

func (s *MyHashSet) Add(key int) {
	begin := s.set
	for begin.next != nil {
		if begin.val == key {
			return
		}
		begin = begin.next
	}
	if begin.val == key {
		return
	}
	begin.next = &List{val: key}
}

func (s *MyHashSet) Remove(key int) {
	begin := s.set
	prev := s.set
	for begin.next != nil {
		if begin.val == key {
			prev.next = begin.next
			return
		}
		prev = begin
		begin = begin.next
	}
	if begin.val == key {
		prev.next = nil
	}
}

func (s *MyHashSet) Contains(key int) bool {
	begin := s.set
	for begin.next != nil {
		if begin.val == key {
			return true
		}
		begin = begin.next
	}
	if begin.val == key {
		return true
	}

	return false
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
	s := getBitString(n)
	fmt.Println(s)
	return s[k-1]
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var iter1 = m - 1
	var iter2 = n - 1

	for i := m + n - 1; i >= 0; i-- {
		if iter1 < 0 {
			nums1[i] = nums2[iter2]
			iter2--
			continue
		} else if iter2 < 0 {
			nums1[i] = nums1[iter1]
			iter1--
			continue
		}

		if nums1[iter1] > nums2[iter2] {
			nums1[i] = nums1[iter1]
			iter1--
		} else {
			nums1[i] = nums2[iter2]
			iter2--
		}
	}
}

func finalPrices_1475(prices []int) []int {
	minArr := func(array []int, value int) int {
		for _, v := range array {
			if v <= value {
				return v
			}
		}
		return value + 1
	}

	for index, value := range prices {
		nextPrice := value + 1
		if index+1 != len(prices) {
			nextPrice = minArr(prices[index+1:], value)
		}
		if nextPrice <= value {
			prices[index] = value - nextPrice
		}
	}

	return prices
}

func main() {

}
