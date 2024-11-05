package linked_list

import (
	"reflect"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		l1 := CreateListNode(1, 2, 4)
		l2 := CreateListNode(1, 3, 4)

		want := []int{1, 1, 2, 3, 4, 4}
		got := MergeTwoSortedLists(l1, l2).Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		l1 := CreateListNode()
		l2 := CreateListNode()

		want := 0
		got := len(MergeTwoSortedLists(l1, l2).Values())

		if want != got {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		l1 := CreateListNode()
		l2 := CreateListNode(0)

		want := []int{0}
		got := MergeTwoSortedLists(l1, l2).Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
