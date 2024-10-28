package maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := 6
		got := MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

		if want != got {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := 1
		got := MaxSubArray([]int{1})

		if want != got {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := 23
		got := MaxSubArray([]int{5, 4, -1, 7, 8})

		if want != got {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
