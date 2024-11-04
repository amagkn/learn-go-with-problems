package valid_anagram

/*
ValidAnagram

Given two strings s and t, return true if t is an
anagram of s, and false otherwise.
*/
func ValidAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := map[rune]int{}

	for _, r := range s {
		dict[r]++
	}

	for _, r := range t {
		if dict[r] == 0 {
			return false
		}

		dict[r]--
	}

	return true
}
