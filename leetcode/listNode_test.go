package leetcode

import (
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
