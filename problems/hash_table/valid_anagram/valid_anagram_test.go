package valid_anagram

import "testing"

func TestValidAnagram(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := true
		got := ValidAnagram("anagram", "nagaram")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := false
		got := ValidAnagram("rat", "car")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := false
		got := ValidAnagram("abcc", "abc")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
