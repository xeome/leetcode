package main

func isSymmetric(root *TreeNode) bool {
	return traverse(root.Left, root.Right)
}

func traverse(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if !compareNode(p, q) {
		return false
	}
	return traverse(p.Left, q.Right) && traverse(p.Right, q.Left)
}
