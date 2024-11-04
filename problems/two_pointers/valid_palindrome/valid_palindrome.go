package valid_palindrome

import (
	"learn-go-with-problems/my_helpers"
	"unicode"
)

/*
ValidPalindrome

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,it reads the same forward and backward. Alphanumeric characters include letters and numbers.
Given a string s, return true if it is a palindrome, or false otherwise.
*/
func ValidPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		runeFromStart := rune(s[l])
		runeFromEnd := rune(s[r])

		if !my_helpers.IsAlphanumericRune(runeFromStart) {
			l++
			continue
		}

		if !my_helpers.IsAlphanumericRune(runeFromEnd) {
			r--
			continue
		}

		if unicode.ToLower(runeFromStart) != unicode.ToLower(runeFromEnd) {
			return false
		}

		l++
		r--
	}

	return true
}
