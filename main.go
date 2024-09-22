package main

import (
	"fmt"
)

func mergeTail(nums1 []int, i int, nums2 []int, n int) {
	for n >= 0 {
		nums1[i] = nums2[n]
		n--
		i--
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	i := len(nums1) - 1
	m = m - 1
	n = n - 1

	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		} else {
			nums1[i] = nums2[n]
			n--
		}
		i--
	}

	if n != 0 {
		mergeTail(nums1, i, nums2, n)
	} else if m != 0 {
		mergeTail(nums1, i, nums1, m)
	}

}

func main() {

	test := []int{1, 2, 3, 0, 0, 0}

	merge(test, 3, []int{2, 5, 6}, 3)

	fmt.Println(test)
}
