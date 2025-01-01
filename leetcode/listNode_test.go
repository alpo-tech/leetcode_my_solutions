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
	list1 := CreateList([]int{1,4,3,2,5,2})
	list1Answer := CreateList([]int{1,2,2,4,3,5})
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
