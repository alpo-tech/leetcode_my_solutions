package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var nums = []int{}
	if root == nil {
		return nums
	}

	nums = append(nums, inorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorderTraversal(root.Right)...)
	return nums
}

func helperInorderTraversalReq(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	helperInorderTraversalReq(root.Left, nums)
	*nums = append(*nums, root.Val)
	helperInorderTraversalReq(root.Right, nums)
}

func inorderTraversalReq(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nums := make([]int, 0)
	helperInorderTraversalReq(root, &nums)
	return nums
}

func InorderTraversalIter(root *TreeNode) []int {
	return []int{}
}

func IsSameTreeHelper(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTreeHelper(p.Right, q.Right) && IsSameTreeHelper(p.Left, q.Left)
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return IsSameTreeHelper(p, q)
}

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftStep := maxDepthHelper(root.Left, 1)
	rightStep := maxDepthHelper(root.Right, 1)
	return int(math.Max(float64(leftStep), float64(rightStep)))
}
