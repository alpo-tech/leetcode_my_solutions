package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{nums: []int{5, 2, 3, 1}}, want: []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
