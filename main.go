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

type LRUCacheNode struct {
	prev *LRUCacheNode
	next *LRUCacheNode
	key  int
	val  int
}

type LRUCache struct {
	head *LRUCacheNode
	tail *LRUCacheNode
	Mp   map[int]*LRUCacheNode
	size int
	cap  int
}

func ConstructorCache(capacity int) LRUCache {
	tmp_head := &LRUCacheNode{prev: nil, next: nil, key: -1, val: -1}
	head := tmp_head
	i := 1
	for i < capacity {
		tmp_head.next = &LRUCacheNode{prev: tmp_head, next: nil, key: -1, val: -1}
		tmp_head = tmp_head.next
		i++
	}
	return LRUCache{head: head, tail: tmp_head, Mp: make(map[int]*LRUCacheNode), size: 0, cap: capacity}
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.Mp[key]
	if !ok {
		return -1
	}
	c.update(node.key)
	return node.val
}

func (c *LRUCache) Put(key int, value int) {
	_, ok := c.Mp[key]
	if ok {
		c.update(key)
		return
	}
	if c.size == c.cap {
		c.remove(c.tail.key)
	}
	c.insert(key, value)
}

func (c *LRUCache) remove(key int) {
	delete(c.Mp, key)
	c.size--
}

func (c *LRUCache) update(key int) {
	if c.head == c.tail {
		return
	}
	node := c.Mp[key]
	if node == c.head {
		return
	}

	if node == c.tail {
		c.PrintCache()
		c.tail = node.prev
		c.tail.next = nil
		node.next = c.head
		c.head.prev = node
		c.head = node
		c.head.prev = nil
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	c.head.prev = node
	node.next = c.head
	node.prev = nil
	c.head = node
}

func (c *LRUCache) insert(key, value int) {
	if c.tail == c.head {
		c.head.val = value
		c.head.key = key
		c.Mp[key] = c.head
		c.size++
		return
	}
	tmp_node := c.tail
	c.tail = tmp_node.prev
	c.tail.next = nil

	tmp_node.prev = nil
	tmp_node.next = c.head
	tmp_node.key = key
	tmp_node.val = value
	c.head.prev = tmp_node
	c.head = tmp_node
	c.size++
	c.Mp[key] = c.head
}

func (c *LRUCache) PrintCache() {
	tmp := c.head
	for tmp != nil {
		fmt.Printf("key -> %d, val -> %d\n", tmp.key, tmp.val)
		tmp = tmp.next
	}
	fmt.Printf("size = %d, cap = %d\n", c.size, c.cap)
	for key, value := range c.Mp {
		fmt.Printf("key = %d, node_key = %d node_val = %d, address %p\n", key, value.key, value.val, value)
	}
	fmt.Println(c.head, c.tail)
}

func main() {
	cache := ConstructorCache(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Get(2)
	//cache.Get(1)
	//test2 := cache.Get(4)
	//test3 := cache.Get(3)
	//test4 := cache.Get(2)
	//fmt.Println(test1, test2, test3, test4)

	cache.PrintCache()

}
