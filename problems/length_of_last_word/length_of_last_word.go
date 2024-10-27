package length_of_last_word

func LengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if length > 0 {
				return length
			}

			continue
		}

		length++
	}

	return length
}
