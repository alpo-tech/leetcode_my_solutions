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

func TestTitleToNumber(t *testing.T) {
	type args struct {
		columntTitle string
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{"A"}, want: 1},    // TODO: Add test cases.
		{args: args{"AB"}, want: 28},  // TODO: Add test cases.
		{args: args{"ZY"}, want: 701}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TitleToNumber(tt.args.columntTitle); got != tt.want {
				t.Errorf("TitleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIsomorphic_205(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{"badc", "fafa"}, want: false},  // TODO: Add test cases.
		{args: args{"egg", "add"}, want: true},     // TODO: Add test cases.
		{args: args{"foo", "bar"}, want: false},    // TODO: Add test cases.
		{args: args{"paper", "title"}, want: true}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsIsomorphic_205(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsIsomorphic_205() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{[]int{1, 2, 1, 4}}, want: true},          // TODO: Add test cases.
		{args: args{[]int{1, 2, 3, 4}}, want: false},         // TODO: Add test cases.
		{args: args{[]int{1, 1, 3, 4, 2, 3, 4}}, want: true}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ContainsDuplicate_217(tt.args.nums); got != tt.want {
				t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsDuplicateNearby(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{[]int{1, 2, 3, 1}, 3}, want: true},        // TODO: Add test cases.
		{args: args{[]int{1, 0, 1, 1}, 1}, want: true},        // TODO: Add test cases.
		{args: args{[]int{1, 2, 3, 1, 2, 3}, 2}, want: false}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ContainsNearbyDuplicate_219(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("ContainsDuplicateNearby() = %v, want %v", got, tt.want)
			}
		})
	}
}
