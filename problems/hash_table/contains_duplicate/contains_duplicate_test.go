package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := true
		got := ContainsDuplicate([]int{1, 2, 3, 1})

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := false
		got := ContainsDuplicate([]int{1, 2, 3, 4})

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := true
		got := ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
