package main

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	// Initialize the head of the linked list
	head := &ListNode{Val: nums[0]}
	// Initialize the current node
	current := head
	// Iterate through the rest of the list
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func createTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// Initialize the root of the tree
	root := &TreeNode{Val: nums[0]}
	// Initialize the queue of nodes, and add the root
	queue := []*TreeNode{root}
	// Iterate through the rest of the list of nodes, and add them to the queue as children of the nodes in the queue (in order)
	for i := 1; i < len(nums); i += 2 {
		// Get the current node from the queue
		current := queue[0]
		// Remove the current node from the queue
		queue = queue[1:]
		// Add the left and right children to the queue, if they exist (i.e. if they are not -1) and add them to the current node as children
		if nums[i] != -1 {
			current.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, current.Left)
		}
		if i+1 < len(nums) && nums[i+1] != -1 {
			current.Right = &TreeNode{Val: nums[i+1]}
			queue = append(queue, current.Right)
		}
	}
	return root
}
