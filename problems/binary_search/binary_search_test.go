package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := 4
		got := BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := -1
		got := BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 2)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
