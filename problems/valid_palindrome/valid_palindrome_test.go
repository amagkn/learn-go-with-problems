package valid_palindrome

import "testing"

func TestValidPalindrome(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := true
		got := ValidPalindrome("A man, a plan, a canal: Panama")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := false
		got := ValidPalindrome("race a car")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := true
		got := ValidPalindrome(" ")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example 4", func(t *testing.T) {
		want := true
		got := ValidPalindrome("Довод")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
