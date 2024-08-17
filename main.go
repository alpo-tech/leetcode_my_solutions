package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	len := len(*h)
	el := (*h)[len-1]
	*h = (*h)[:len]
	return el
}

type KthLargest struct {
	myHeap *IntHeap
	k      int
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	kthLargest := KthLargest{k: k, myHeap: minHeap}
	for _, num := range nums {
		kthLargest.Add(num)
	}
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	if this.myHeap.Len() < this.k {
		heap.Push(this.myHeap, val)
	} else if val > (*this.myHeap)[0] {
		heap.Pop(this.myHeap)
		heap.Push(this.myHeap, val)
	}

	return (*this.myHeap)[0]

}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	heap.Init()
}
