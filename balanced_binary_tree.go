package main

import "math"

func isBalanced(root *TreeNode) bool {
	_, ok := dfs(root)
	return ok
}

// dfs is a depth-first search function that calculates the height of each subtree
// rooted at node and returns the height and a boolean indicating whether the
// subtree is balanced.
//
// Parameters:
//   - node: The root node of the subtree to check.
//
// Returns:
//
//	A tuple containing the height of the subtree and a boolean indicating
//	whether the subtree is balanced.
func dfs(node *TreeNode) (int, bool) {
	// If the node is nil, the subtree is balanced and has height 0
	if node == nil {
		return 0, true
	}
	// Recursively call dfs on the left and right subtrees of the node
	// and store the returned heights and boolean values in variables
	left, leftBool := dfs(node.Left)
	right, rightBool := dfs(node.Right)
	// Calculate the height of the subtree rooted at node as 1 plus the maximum
	// of the heights of the left and right subtrees
	height := 1 + max(left, right)
	// Check if the absolute difference between the heights of the left and right
	// subtrees is less than 2, and if both the left and right subtrees are balanced.
	// If both conditions are true, the subtree rooted at node is balanced and return
	// height and true. Otherwise, return height and false.
	return height, int(math.Abs(float64(left-right))) < 2 && leftBool && rightBool
}
