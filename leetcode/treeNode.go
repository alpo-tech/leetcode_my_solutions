package leetcode

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
