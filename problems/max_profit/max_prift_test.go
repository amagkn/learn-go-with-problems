package max_profit

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := 5
		got := MaxProfit([]int{7, 1, 5, 3, 6, 4})

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := 0
		got := MaxProfit([]int{7, 6, 4, 3, 1})

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
