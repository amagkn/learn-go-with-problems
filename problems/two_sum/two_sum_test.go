package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := []int{0, 1}
		got := TwoSum([]int{2, 7, 11, 15}, 9)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := []int{1, 2}
		got := TwoSum([]int{3, 2, 4}, 6)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := []int{0, 1}
		got := TwoSum([]int{3, 3}, 6)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
