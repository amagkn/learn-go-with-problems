package linked_list

import (
	"reflect"
	"testing"
)

func TestListNode(t *testing.T) {
	t.Run("Values of list node", func(t *testing.T) {
		linkedList := CreateListNode(1, 2, 3, 4, 5)

		want := []int{1, 2, 3, 4, 5}
		got := linkedList.Values()

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
