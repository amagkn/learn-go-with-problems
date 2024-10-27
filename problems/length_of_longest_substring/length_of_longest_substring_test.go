package length_of_longest_substring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := 3
		got := LengthOfLongestSubstring("abcabcbb")

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := 1
		got := LengthOfLongestSubstring("bbbbb")

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := 3
		got := LengthOfLongestSubstring("pwwkew")

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
