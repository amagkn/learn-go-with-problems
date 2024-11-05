package linked_list

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
