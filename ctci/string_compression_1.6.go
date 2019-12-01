package ctci

import (
	"strconv"
	"strings"
)

// compress compresses a string using counts of repeated characters
// if the compressed version of the string is not smaller than the original, this function will return the original string
// e.g. "aaabcccccaaa" becomes "a3b1c5a3"
func compress(s string) string {
	if len(s) <= 2 { // The compressed string will not become smaller than the original
		return s
	}
	var sb strings.Builder
	prevChar := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == prevChar {
			count++
		} else {
			sb.WriteString(string(prevChar) + strconv.Itoa(count))
			count = 1
			prevChar = s[i]
		}
	}
	sb.WriteString(string(prevChar) + strconv.Itoa(count))
	ret := sb.String()
	if len(ret) < len(s) {
		return ret
	}
	return s
}
