package linked_list

/*
MergeTwoSortedLists

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.
*/
func MergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}

	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}

		curr = curr.Next
	}

	if l1 != nil {
		curr.Next = l1
	}

	if l2 != nil {
		curr.Next = l2
	}

	return dummy.Next
}
