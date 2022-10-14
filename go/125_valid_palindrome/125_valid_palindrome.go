package leetcode

import "unicode"

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) {
			r--
			continue
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		r--
		l++
	}
	return true
}
