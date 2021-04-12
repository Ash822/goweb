package msgutils

import "strings"

func Reverse(str string) string {
	revered := ""
	for _, c := range str {
		revered = string(c) + revered
	}
	return revered
}

func IsPalindrome(str string) bool {
	if str == "" {
		return false
	}

	str = strings.ToLower(str)

	if str == Reverse(str) {
		return true
	}
	return false
}
