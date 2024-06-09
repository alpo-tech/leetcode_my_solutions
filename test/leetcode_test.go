package test

import (
	"leetcode.com/leetcode"
	"reflect"
	"testing"
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
