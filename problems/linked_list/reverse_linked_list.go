package linked_list

func ReverseLinkedList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head

	for curr != nil {
		next := curr.Next

		curr.Next = prev
		prev = curr

		curr = next
	}

	return prev
}
