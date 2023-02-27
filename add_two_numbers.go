package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a new node to store the result
	result := &ListNode{}
	// Create a pointer to the result node
	p := result
	// Create a carry variable to store the carry
	carry := 0
	// Iterate over the two lists, O(n)
	for l1 != nil || l2 != nil {
		// If the list is not nil, add the value to the carry
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		// If the list is not nil, add the value to the carry
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		// Create a new node with the carry % 10 as the value
		p.Next = &ListNode{Val: carry % 10}
		// Move the pointer to the next node
		p = p.Next
		// Update the carry
		carry /= 10
	}
	// If the carry is not 0, create a new node with the carry as the value
	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}
	// Return the next node of the result node
	return result.Next

}
