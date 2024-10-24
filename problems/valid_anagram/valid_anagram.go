package valid_anagram

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
