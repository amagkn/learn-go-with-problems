package my_helpers

import (
	"regexp"
)

var re, _ = regexp.Compile("^[a-zA-Zа-яА-Я0-9]$")

/*
IsAlphanumericRune

Проверяет символ является ли он буквой или числом.
В реализации используется регулярка, но существует решение с помощью пакета unicode (unicode.IsLetter || unicode.IsDigit)
*/
func IsAlphanumericRune(r rune) bool {
	return re.MatchString(string(r))
}
