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

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"world      ", 5},
		{"         world", 5},
		{"world", 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.LenghtOfLastWord(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LengOfLastWord(%v) = %v; want = %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.PlusOne(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusOne(%v) = %v; want = %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.AddBinary(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBinary(%v, %v) = %v; want = %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMySqrt(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{4, 2},
		{8, 2},
		{0, 0},
		{1, 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.MySqrt(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySqrt(%v) = %v; want = %v", tt.want, got, tt.want)
			}
		})
	}
}

func TestClimbingsStairs(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.ClimbingsStairsReqursive(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClimbingStartsReq(%v) = %v; want = %v", tt.want, got, tt.want)
			}

			got = leetcode.ClimbingsStairsIter(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClimbingStartsIter(%v) = %v; want = %v", tt.want, got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicateFromSortedList(t *testing.T) {

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
			leetcode.MergeSortedArray(tt.array1, tt.count1, tt.array2, tt.count2)
			if !reflect.DeepEqual(tt.array1, tt.want) {
				t.Errorf("MergeSortedArray(%v, %v, %v, %v) = %v; want = %v", tt.array1, tt.count1, tt.array2, tt.count2, tt.array1, tt.want)
			}
		})
	}
}

func TestReverseParenthese(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"(abcd)", "dcba"},
		{"(u(love)i)", "iloveu"},
		{"(ed(et(oc))el)", "leetcode"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.ReverseParentheses(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseParenthese(%v) = %v; want = %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestInorderTraversal(t *testing.T) {
	input1 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(1),
		nil,
		leetcode.IntPtr(2),
		leetcode.IntPtr(3),
	})

	input2 := leetcode.CreateTreeNode([]*int{})

	input3 := leetcode.CreateTreeNode([]*int{leetcode.IntPtr(1)})

	tests := []struct {
		input *leetcode.TreeNode
		want  []int
	}{
		{input1, []int{1, 3, 2, 4}},
		{input2, []int{}},
		{input3, []int{1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.InorderTraversalReq(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InorderTraversalReq(%v) = %v; want = %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaxDepth(t *testing.T) {
	input1 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(3),
		leetcode.IntPtr(9),
		leetcode.IntPtr(20),
		nil,
		nil,
		leetcode.IntPtr(15),
		leetcode.IntPtr(7),
	})

	input2 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(1),
		nil,
		leetcode.IntPtr(2),
	})

	tests := []struct {
		input *leetcode.TreeNode
		want  int
	}{
		{input1, 3},
		{input2, 2},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.MaxDepth(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxDepth(%v) = %v, want = %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestSortedArrayToBST(t *testing.T) {

}

func TestIsBalanced(t *testing.T) {
	input1 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(1),
		leetcode.IntPtr(2),
		leetcode.IntPtr(2),
		leetcode.IntPtr(3),
		nil,
		nil,
		leetcode.IntPtr(3),
		leetcode.IntPtr(4),
		nil,
		nil,
		leetcode.IntPtr(4),
	})

	tests := []struct {
		input *leetcode.TreeNode
		want  bool
	}{
		{input1, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.IsBalanced(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsBalanced(%v)=%v want=%v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMinDepth(t *testing.T) {
	input1 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(3),
		leetcode.IntPtr(9),
		leetcode.IntPtr(20),
		nil,
		nil,
		leetcode.IntPtr(15),
		leetcode.IntPtr(7),
	})

	input2 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(2),
		nil,
		leetcode.IntPtr(3),
		nil,
		leetcode.IntPtr(4),
		nil,
		leetcode.IntPtr(5),
		nil,
		leetcode.IntPtr(6),
	})

	tests := []struct {
		input *leetcode.TreeNode
		want  int
	}{
		{input1, 2},
		{input2, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.MinDepth(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinDepth(%v)=%v want %v", tt.input, got, tt.want)
			}
		})
	}

}

func TestHasPathSum(t *testing.T) {
	input1 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(5),
		leetcode.IntPtr(4),
		leetcode.IntPtr(8),
		leetcode.IntPtr(11),
		nil,
		leetcode.IntPtr(13),
		leetcode.IntPtr(4),
		leetcode.IntPtr(7),
		leetcode.IntPtr(2),
		nil,
		nil,
		nil,
		leetcode.IntPtr(1),
	})
	input2 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(1),
		leetcode.IntPtr(2),
		leetcode.IntPtr(3),
	})

	input3 := leetcode.CreateTreeNode([]*int{
		leetcode.IntPtr(1),
		leetcode.IntPtr(-2),
		leetcode.IntPtr(-3),
		leetcode.IntPtr(1),
		leetcode.IntPtr(3),
		leetcode.IntPtr(-2),
		nil,
		leetcode.IntPtr(-1),
	})

	tests := []struct {
		input     *leetcode.TreeNode
		targetSum int
		want      bool
	}{
		{input1, 22, true},
		{input2, 5, false},
		{input3, -1, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.HasPathSum(tt.input, tt.targetSum)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasPathSum(%v, %v) = %v, want = %v", tt.input, tt.targetSum, got, tt.want)
			}
		})
	}
}

func TestGenerageTrianglePascal(t *testing.T) {
	tests := []struct {
		input int
		want  [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.GenerageTrianglePascal(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerageTrianglePascal(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestGetRowTriangelPascal(t *testing.T) {
	tests := []struct {
		input int
		want  []int
	}{
		{3, []int{1, 3, 3, 1}},
		{0, []int{1}},
		{1, []int{1, 1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.GetRowTriangelPascalReqursive(tt.input)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("GetRowTriangelPascal(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.MaxProfixBruteForce_121(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxProfitBruteForce(%v)=%v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindromeString(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"Marge, let's \"[went].\" I await {news} telegram.", true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := leetcode.IsPalindromeString_125(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsPalindromeString(%v) = %v ; want %v", tt.input, got, tt.input)
			}
		})
	}
}
