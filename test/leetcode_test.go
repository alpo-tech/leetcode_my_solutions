package test

import (
	"reflect"
	"testing"

	"leetcode.com/leetcode"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3}, 7, []int{0, 0}}, // тест на случай, когда нет решения
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{121, true},
		{-123, false},
		{10, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.IsPalindromeV2(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsPalindrome(%d) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.RomanToInt(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RomanToInt(%s) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"caca", "c", "cac"}, "c"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.LongestCommonPrefix(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LongestCommonPrefix(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"(]", false},
		{"()", true},
		{"()[]{}", true},
		{"(({{}}))", true},
		{"(])", false},
		{"((([)))", false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.ValidParenthesesV2(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidParentheses(%s) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input       []int
		want        int
		updateInput []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.RemoveDuplicatess(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates(%v) = %v; want %v", tt.input, got, tt.want)
			}

		})
	}
}

func TestRemoveELement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.RemoveElement(tt.nums, tt.val)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElemnet(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		str    string
		substr string
		want   int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"abcabd", "abcabd", 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.StrStr(tt.str, tt.substr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrStr(%v. %v) = %v; want %v", tt.str, tt.substr, got, tt.want)
			}
		})
	}
}

func TestSearchInsertPosition(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1}, 2, 1},
		{[]int{1}, 0, 0},

		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.SearchInsertPosition(tt.input, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchInsertPosition(%v, %v) = %v; want %v", tt.input, tt.target, got, tt.want)
			}
		})
	}
}
