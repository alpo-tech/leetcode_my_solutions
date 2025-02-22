package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintTree(root *TreeNode) {
	arr := make([]int, 0)
	var helperPrint func(*TreeNode, *[]int)
	helperPrint = func(root *TreeNode, array *[]int) {
		*array = append(*array, root.Val)

		if root.Left == nil && root.Right == nil {
			return
		}

		if root.Left != nil {
			helperPrint(root.Left, array)
		} else {
			*array = append(*array, -1)
		}
		if root.Right != nil {
			helperPrint(root.Right, array)
		} else {
			*array = append(*array, -1)
		}
	}
	if root != nil {
		helperPrint(root, &arr)
	}

	fmt.Println(arr)
}

func CreateTreeNode(array []*int) *TreeNode {
	if len(array) == 0 || array[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *array[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(array) {
		current := queue[0]
		queue = queue[1:]

		if i < len(array) && array[i] != nil {
			current.Left = &TreeNode{Val: *array[i]}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(array) && array[i] != nil {
			current.Right = &TreeNode{Val: *array[i]}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}

func IntPtr(i int) *int {
	return &i
}

func helperInorderTraversalReq(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	helperInorderTraversalReq(root.Left, nums)
	*nums = append(*nums, root.Val)
	helperInorderTraversalReq(root.Right, nums)
}

func InorderTraversalReq(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nums := make([]int, 0)
	helperInorderTraversalReq(root, &nums)
	return nums
}

// TODO: реализовать прямоц обход дерева с помощью итерации
func InorderTraversalIter(root *TreeNode) []int {
	return []int{}
}

func IsSameHelper(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameHelper(p.Right, q.Right) && IsSameHelper(p.Left, q.Left)
}

func IsSame(p *TreeNode, q *TreeNode) bool {
	return IsSameHelper(p, q)
}

// this is test

func IsSymmetricHelper(right *TreeNode, left *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if right == nil || left == nil {
		return false
	}

	if right.Val != left.Val {
		return false
	}

	return IsSymmetricHelper(right.Right, left.Left) && IsSymmetricHelper(right.Left, left.Right)
}

func isSymmetric(root *TreeNode) bool {
	return IsSymmetricHelper(root.Right, root.Left)
}

func maxDepthHelper(root *TreeNode, step int) int {
	if root == nil {
		return step
	}

	return int(math.Max(float64(maxDepthHelper(root.Right, step+1)), float64(maxDepthHelper(root.Left, step+1))))
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftStep := maxDepthHelper(root.Left, 1)
	rightStep := maxDepthHelper(root.Right, 1)
	return int(math.Max(float64(leftStep), float64(rightStep)))
}

func SortedArrayToBst(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	root := &TreeNode{Val: nums[middle]}
	root.Left = SortedArrayToBst(nums[:middle])
	root.Right = SortedArrayToBst(nums[middle+1:])
	return root
}

func dfs(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := dfs(root.Left)
	isRightBalanced, rightHeight := dfs(root.Right)

	diff := math.Abs(float64(leftHeight - rightHeight))
	if isLeftBalanced && isRightBalanced && diff <= 1 {
		return true, 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
	}
	return false, -1
}

func IsBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}

func MinDepthHelper(root *TreeNode, step int) (int, bool) {
	if root == nil {
		return step, false
	}

	if root.Left == nil && root.Right == nil {
		return step + 1, true
	}

	minLeft, okLeft := MinDepthHelper(root.Left, step+1)
	minRight, okRight := MinDepthHelper(root.Right, step+1)

	if okLeft && okRight {
		return min(minLeft, minRight), true
	}

	if okLeft {
		return minLeft, true
	}

	return minRight, true

}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minLeft, okLeft := MinDepthHelper(root.Left, 1)
	minRight, okRight := MinDepthHelper(root.Right, 1)

	if okLeft && okRight {
		return min(minLeft, minRight)
	}

	if okLeft {
		return minLeft
	}

	return minRight

}

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && targetSum-root.Val == 0 {
		return true
	}

	return HasPathSum(root.Left, targetSum-root.Val) || HasPathSum(root.Right, targetSum-root.Val)
}

func helperPreorderTraversalReq(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	helperInorderTraversalReq(root.Left, nums)
	helperInorderTraversalReq(root.Right, nums)
}

func PreorderTraversalReq(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nums := make([]int, 0)
	helperPreorderTraversalReq(root, &nums)
	return nums
}

func postorderHelper(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	postorderHelper(root.Right, nums)
	postorderHelper(root.Left, nums)
	*nums = append(*nums, root.Val)
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := make([]int, 0)
	postorderHelper(root, &nums)
	return nums
}

func CountNodes_222(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + CountNodes_222(root.Left) + CountNodes_222(root.Right)
}

func BinaryTreePathsHelper(root *TreeNode, str string) []string {

	commonStr := str + "->" + strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		return []string{commonStr}
	}

	if root.Left == nil {
		return BinaryTreePathsHelper(root.Right, commonStr)
	}

	if root.Right == nil {
		return BinaryTreePathsHelper(root.Left, commonStr)
	}

	return append(BinaryTreePathsHelper(root.Right, commonStr), BinaryTreePathsHelper(root.Left, commonStr)...)

}

func BinaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return []string{}
	}

	result := make([]string, 0)

	if root.Left == nil && root.Right == nil {
		return append(result, strconv.Itoa(root.Val))
	}

	if root.Left != nil {
		result = append(result, BinaryTreePathsHelper(root.Left, strconv.Itoa(root.Val))...)
	}

	if root.Right != nil {
		result = append(result, BinaryTreePathsHelper(root.Right, strconv.Itoa(root.Val))...)
	}

	return result
}

func Flatten(root *TreeNode) {
}
