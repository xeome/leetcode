package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
