package linked_list

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		l1 := CreateListNode(2, 4, 3)
		l2 := CreateListNode(5, 6, 4)

		want := []int{7, 0, 8} // Explanation: 342 + 465 = 807
		got := AddTwoNumbers(l1, l2).Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		l1 := CreateListNode(0)
		l2 := CreateListNode(0)

		want := []int{0}
		got := AddTwoNumbers(l1, l2).Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		l1 := CreateListNode(9, 9, 9, 9, 9, 9, 9)
		l2 := CreateListNode(9, 9, 9, 9)

		want := []int{8, 9, 9, 9, 0, 0, 0, 1}
		got := AddTwoNumbers(l1, l2).Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
