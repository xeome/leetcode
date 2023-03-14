package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	if root.Left == nil {
		return hasPathSum(root.Right, targetSum-root.Val)
	}
	if root.Right == nil {
		return hasPathSum(root.Left, targetSum-root.Val)
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
