package leetcode

import (
	"reflect"
	"testing"
)

func TestConvertToTitle_168(t *testing.T) {
	type args struct {
		columnNumber int
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{1}, want: "A"},    // TODO: Add test cases.
		{args: args{27}, want: "AA"},  // TODO: Add test cases.
		{args: args{28}, want: "AB"},  // TODO: Add test cases.
		{args: args{701}, want: "ZY"}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ConvertToTitle_168(tt.args.columnNumber); got != tt.want {
				t.Errorf("ConvertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMajorityElement_169(t *testing.T) {
	tests := []struct {
		args []int
		want int
	}{
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := MajorityElement_169(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MajorityElement(%v)=%v, want %v", tt.args, got, tt.want)
			}
		})

		t.Run("version algo", func(t *testing.T) {
			got := MajorityElement_VoteAlgorithm_169(tt.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MajorityElementAlg(%v)=%v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
