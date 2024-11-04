package top_k_frequent_elements

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := []int{1, 2}
		got := TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := []int{1}
		got := TopKFrequent([]int{1}, 1)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
