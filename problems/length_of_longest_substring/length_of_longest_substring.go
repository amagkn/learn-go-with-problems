package length_of_longest_substring

/*
LengthOfLongestSubstring

Given a string s, find the length of the longest
substring without repeating characters.
*/
func LengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)

	maxLength := 0
	start := 0

	for end, char := range s {
		if lastIndex, ok := charIndex[char]; ok && lastIndex >= start {
			start = lastIndex + 1
		}

		charIndex[char] = end

		currentLength := end - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
