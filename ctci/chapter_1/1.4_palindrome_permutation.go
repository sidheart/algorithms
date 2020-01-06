package chapter_1

import "unicode"

func isPalindromePermutation(s string) bool {
	charFreq := make(map[rune]int)
	for _, c := range s {
		if c != ' ' {
			charFreq[unicode.ToLower(c)] += 1
		}
	}
	var oddFrequencyExists bool
	for _, freq := range charFreq {
		if freq % 2 != 0 {
			if !oddFrequencyExists {
				oddFrequencyExists = true
				continue
			}
			return false
		}
	}
	return true
}
