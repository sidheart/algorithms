package ctci

import "unicode"

func isUnique(s string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		if _, ok := m[c]; ok {
			return false
		}
		m[c]++
	}
	return true
}

func toIndex(r rune) rune {
	if unicode.IsUpper(r) {
		return r - 'A'
	}
	return r - 'a' + 26
}

func isUniqueAscii(s string) bool {
	var ascii [42]int
	for _, r := range s {
		ascii[toIndex(r)]++
		if ascii[toIndex(r)] > 1 {
			return false
		}
	}
	return true
}
