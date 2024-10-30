package leetcode

import (
	"reflect"
	"testing"
)

func Test_findKthBit(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthBit(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBits_338(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{2}, want: []int{0, 1, 1}},
		{args: args{5}, want: []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countBits_338(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits_338() = %v, want %v", got, tt.want)
			}
		})
	}
}
