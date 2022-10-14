package leetcode

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}
	a := 'a'
	for i := 0; i <= len(palindrome)/2-1; i++ {
		if int(palindrome[i]) <= int(a) {
			continue
		} else {
			return palindrome[:i] + string(a) + palindrome[i+1:]
		}
	}
	return palindrome[:len(palindrome)-1] + string("b")
}
