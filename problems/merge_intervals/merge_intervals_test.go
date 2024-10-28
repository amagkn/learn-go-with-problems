package merge_intervals

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := [][]int{{1, 6}, {8, 10}, {15, 18}}
		got := MergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := [][]int{{1, 5}}
		got := MergeIntervals([][]int{{1, 4}, {4, 5}})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
