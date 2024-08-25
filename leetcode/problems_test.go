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

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{"anagram", "nagaram"}, want: true},
		{args: args{"rat", "car"}, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{38}, want: 2},
		{args: args{0}, want: 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := AddDigits(tt.args.num); got != tt.want {
				t.Errorf("AddDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanNimWim_292(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{4}, false},
		{args{1}, true},
		{args{2}, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CanNimWim_292(tt.args.n); got != tt.want {
				t.Errorf("CanNimWim_292() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectCapitalUse_520(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{"USA"}, true},
		{args{"FlaG"}, false},
		{args{"leetcode"}, true},
		{args{"Google"}, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := DetectCapitalUse_520(tt.args.word); got != tt.want {
				t.Errorf("DetectCapitalUse_520() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfit_122(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{[]int{7, 1, 5, 3, 6, 4}}, want: 7},
		{args: args{[]int{1, 2, 3, 4, 5}}, want: 4},
		{args: args{[]int{7, 6, 4, 3, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxProfit_122(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit_122() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSubsequence_392(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{"acb", "ahbgdc"}, want: false},
		{args: args{"abc", "ahbgdc"}, want: true},
		{args: args{"axc", "ahbgdc"}, want: false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsSubsequence_392(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsSubsequence_392() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanConstruct_383(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{"a", "b"}, want: false},
		{args: args{"aa", "ab"}, want: false},
		{args: args{"aa", "aab"}, want: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CanConstruct_383(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("CanConstruct_383() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordPattern_290(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args{"abba", "dog cat cat dog"}, true},
		{args{"abba", "dog cat cat fish"}, false},
		{args{"aaaa", "dog cat cat dog"}, false},
		{args{"abba", "dog dog dog dog"}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := WordPattern_290(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("WordPattern_290() = %v, want %v", got, tt.want)
			}
		})
	}
}
