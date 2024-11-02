package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(values ...int) *ListNode {
	dummy := &ListNode{}

	curr := dummy

	for _, v := range values {
		nextListNode := &ListNode{v, nil}

		curr.Next = nextListNode
		curr = curr.Next
	}

	return dummy.Next
}

func (l *ListNode) Values() []int {
	var values []int

	curr := l

	for curr != nil {
		values = append(values, curr.Val)

		curr = curr.Next
	}

	return values
}
