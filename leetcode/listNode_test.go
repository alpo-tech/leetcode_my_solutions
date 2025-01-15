package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseBetween_92(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ReverseBetween_92(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseBetween_92() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDecimalValue_1290(t *testing.T) {
	type args struct {
		head *ListNode
	}
	list1 := CreateList([]int{1, 0, 1})
	list2 := CreateList([]int{0})
	tests := []struct {
		args args
		want int
	}{
		{args: args{list1}, want: 5},
		{args: args{list2}, want: 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getDecimalValue_1290(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue_1290() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	list1 := CreateList([]int{1, 2, 3, 4, 5})
	list1Answer := list1.Next.Next

	list2 := CreateList([]int{1, 2, 3, 4, 5, 6})
	list2Answer := list2.Next.Next.Next
	PrintList(list1)
	fmt.Println()
	PrintList(list1Answer)

	tests := []struct {
		args args
		want *ListNode
	}{
		{args: args{list1}, want: list1Answer},
		{args: args{list2}, want: list2Answer},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := middleNode_876(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition_86(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	list1 := CreateList([]int{1, 4, 3, 2, 5, 2})
	list1Answer := CreateList([]int{1, 2, 2, 4, 3, 5})
	tests := []struct {
		args args
		want *ListNode
	}{
		{args: args{list1, 3}, want: list1Answer},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Partition_86(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Partition_86() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteNode_237(t *testing.T) {
	type args struct {
		node *ListNode
	}
	list1 := CreateList([]int{4, 5, 1, 9})
	list1Answer := CreateList([]int{4, 1, 9})

	list2 := CreateList([]int{4, 5, 1, 9})
	list2Answer := CreateList([]int{4, 5, 9})

	tests := []struct {
		args args
		want *ListNode
	}{
		{args: args{list1}, want: list1Answer},
		{args: args{list2}, want: list2Answer},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			deleteNode_237(tt.args.node)
		})
	}
}

func Test_modifiedList_3217(t *testing.T) {
	type args struct {
		nums []int
		head *ListNode
	}

	list1 := CreateList([]int{1, 2, 3, 4, 5})
	list1Answer := CreateList([]int{4, 5})

	list2 := CreateList([]int{1, 2, 1, 2, 1, 2})
	list2Answer := CreateList([]int{2, 2, 2})

	list3 := CreateList([]int{1, 2, 3, 4})
	list3Answer := CreateList([]int{1, 2, 3, 4})

	list4 := CreateList([]int{2, 10, 9})
	list4Answer := CreateList([]int{10})

	tests := []struct {
		args args
		want *ListNode
	}{
		{args: args{nums: []int{1, 2, 3}, head: list1}, want: list1Answer},
		{args: args{nums: []int{1}, head: list2}, want: list2Answer},
		{args: args{nums: []int{5}, head: list3}, want: list3Answer},
		{args: args{nums: []int{9, 2, 5}, head: list4}, want: list4Answer},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := modifiedList_3217(tt.args.nums, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifiedList_3217() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNodes_2487(t *testing.T) {
	type args struct {
		head *ListNode
	}
	list1 := CreateList([]int{5, 2, 13, 3, 8})
	list1Answer := CreateList([]int{13, 8})

	list2 := CreateList([]int{1, 1, 1, 1})
	list2Answer := CreateList([]int{1, 1, 1, 1})
	tests := []struct {
		args args
		want *ListNode
	}{
		{args: args{head: list1}, want: list1Answer},
		{args: args{head: list2}, want: list2Answer},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := removeNodes_2487(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNodes_2487() = %v, want %v", got, tt.want)
			}
		})
	}
}
