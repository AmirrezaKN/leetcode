package validpalindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")

	right := len(s) - 1

	for left := 0; left < len(s) && left <= right; left++ {
		if s[left] != s[right] {
			return false
		}
		right--
	}

	return true
}
