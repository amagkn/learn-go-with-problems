package group_anagrams

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}}
		got := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := [][]string{{""}}
		got := GroupAnagrams([]string{""})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := [][]string{{"a"}}
		got := GroupAnagrams([]string{"a"})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
