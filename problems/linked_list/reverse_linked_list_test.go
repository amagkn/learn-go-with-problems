package linked_list

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		linkedList := CreateListNode(1, 2, 3, 4, 5)

		want := []int{5, 4, 3, 2, 1}
		got := ReverseLinkedList(linkedList).Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		linkedList := CreateListNode(1, 2)

		want := []int{2, 1}
		got := ReverseLinkedList(linkedList).Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3: Empty linked list", func(t *testing.T) {
		linkedList := CreateListNode()

		want := 0
		got := len(ReverseLinkedList(linkedList).Values())

		if want != got {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
