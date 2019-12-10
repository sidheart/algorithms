package ctci

import "strings"

// isSubstring returns true iff s2 is a substring of s1
func isSubstring(s1, s2 string) bool {
	return strings.Contains(s1, s2)
}

// isRotation returns true iff s2 is a rotation of s1
// e.g. s1 = "abcde" s2 = "eabcd" would return true
// 		s1 = "abcde" s2 = "abecd" would return false
func isRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	i := 0
	for j := 0; j < len(s1); j++ {
		if s1[j] == s2[i] {
			i++
		}
	}
	return isSubstring(s1, s2[i:])
}
