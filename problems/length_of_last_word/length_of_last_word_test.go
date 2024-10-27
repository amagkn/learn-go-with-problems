package length_of_last_word

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		want := 5
		got := LengthOfLastWord("Hello World")

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		want := 4
		got := LengthOfLastWord("   fly me   to   the moon  ")

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		want := 6
		got := LengthOfLastWord("luffy is still joyboy")

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
