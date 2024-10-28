package bubble_sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := []int{0, 0, 1, 1, 2, 2}
		got := []int{2, 0, 2, 1, 1, 0}

		BubbleSort(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := []int{0, 1, 2}
		got := []int{2, 0, 1}

		BubbleSort(got)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
