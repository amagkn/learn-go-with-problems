package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSumWithSortedArray(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := []int{1, 2}
		got := TwoSumWithSortedArray([]int{2, 7, 11, 15}, 9)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := []int{1, 3}
		got := TwoSumWithSortedArray([]int{2, 3, 4}, 6)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := []int{1, 2}
		got := TwoSumWithSortedArray([]int{-1, 0}, -1)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
