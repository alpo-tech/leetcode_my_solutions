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

func TestSummaryRanges_228(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []string
	}{
		{args: args{[]int{0, 1, 2, 4, 5, 7}}, want: []string{"0->2", "4->5", "7"}},
		{args: args{[]int{0, 2, 3, 4, 6, 8, 9}}, want: []string{"0", "2->4", "6", "8->9"}},
		{args: args{[]int{0}}, want: []string{"0"}},
		{args: args{[]int{0, 1, 2, 3, 4}}, want: []string{"0->4"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SummaryRanges_228(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SummaryRanges_228() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHammingWeight_191(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{args{11}, 3},
		{args{128}, 1},
		{args{2147483645}, 30},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := HammingWeight_191(tt.args.n); got != tt.want {
				t.Errorf("HammingWeight_191() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, want: 49},
		{args: args{[]int{1, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxArea(tt.args.height); got != tt.want {
				t.Errorf("MaxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		args args
	}{
		{args: args{[][]int{{1, 1}, {1, 0}}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			GameOfLife(tt.args.board)
		})
	}
}

func Test_spiralOrderHelper(t *testing.T) {
	type args struct {
		matrix      [][]int
		resultOrder []int
		count       int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrderHelper(tt.args.matrix, tt.args.resultOrder, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrderHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{args: args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		array1 []int
		count1 int
		array2 []int
		count2 int
		want   []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			MergeSortedArray_88(tt.array1, tt.count1, tt.array2, tt.count2)
			if !reflect.DeepEqual(tt.array1, tt.want) {
				t.Errorf("MergeSortedArray(%v, %v, %v, %v) = %v; want = %v", tt.array1, tt.count1, tt.array2, tt.count2, tt.array1, tt.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{nums: []int{3, 2, 2, 3}, val: 3}, want: 2},
		{args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2}, want: 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RemoveElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
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
			got := RemoveDuplicatess_26(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates(%v) = %v; want %v", tt.input, got, tt.want)
			}

		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveDuplicatess_26([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	}
}

func Test_canJump_55(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{nums: []int{2, 3, 1, 1, 4}}, want: true},
		{name: "test2", args: args{nums: []int{3, 2, 1, 0, 4}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump_55(tt.args.nums); got != tt.want {
				t.Errorf("canJump_55() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSeniors(t *testing.T) {
	type args struct {
		details []string
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{details: []string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}}, want: 2},
		{args: args{details: []string{"1313579440F2036", "2921522980M5644"}}, want: 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countSeniors(tt.args.details); got != tt.want {
				t.Errorf("countSeniors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoOutOfThree_2032(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		nums3 []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{nums1: []int{1, 1, 3, 2}, nums2: []int{2, 3}, nums3: []int{3}}, want: []int{3, 2}},
		{args: args{nums1: []int{1, 2, 2}, nums2: []int{4, 3, 3}, nums3: []int{5}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := twoOutOfThree_2032(tt.args.nums1, tt.args.nums2, tt.args.nums3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoOutOfThree_2032() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		args args
		want []string
	}{
		{args: args{"this apple is sweet", "this apple is sour"}, want: []string{"sweet", "sour"}},
		{args: args{"apple apple", "banana"}, want: []string{"sweet", "sour"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedSquares_977(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{[]int{-4, -1, 0, 3, 10}}, want: []int{0, 1, 9, 16, 100}},
		{args: args{[]int{-7, -3, 2, 3, 11}}, want: []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := sortedSquares_977(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares_977() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFinalState_3264(t *testing.T) {
	type args struct {
		nums       []int
		k          int
		multiplier int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{[]int{2, 1, 3, 5, 6}, 5, 2}, want: []int{8, 4, 6, 5, 6}},
		{args: args{[]int{1, 2}, 3, 4}, want: []int{16, 8}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getFinalState_3264(tt.args.nums, tt.args.k, tt.args.multiplier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFinalState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDuplicate_287(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{[]int{1, 3, 4, 2, 2}}, want: 2},
		{args: args{[]int{3, 1, 3, 4, 2}}, want: 3},
		{args: args{[]int{3, 3, 3, 3, 3}}, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindDuplicate_287(tt.args.nums); got != tt.want {
				t.Errorf("FindDuplicate_287() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMissingNumber_268(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{[]int{3, 0, 1}}, want: 2},
		{args: args{[]int{0, 1}}, want: 2},
		{args: args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, want: 8},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MissingNumber_268(tt.args.nums); got != tt.want {
				t.Errorf("MissingNumber_268() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findErrorNums_645(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{[]int{1, 2, 2, 4}}, want: []int{2, 3}},
		{args: args{[]int{1, 1}}, want: []int{1, 2}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findErrorNums_645(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums_645() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveZeroes_283(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		// TODOname: Add test cases.
		{args: args{[]int{0, 1, 0, 3, 12}}, want: []int{1, 3, 12, 0, 0}},
		{args: args{[]int{0}}, want: []int{0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			moveZeroes_283(tt.args.nums)
		})
	}
}
