package leetcode

func canMakeWord(word string, letterCounts map[byte]int) bool {
	wordLetterCounts := make(map[byte]int)
	for i, _ := range word {
		wordLetterCounts[word[i]]++
	}
	for k, v := range wordLetterCounts {
		if v > letterCounts[k] {
			return false
		}
	}
	return true
}

func computeWordScore(word string, scoreList []int) (score int) {
	for i, _ := range word {
		score += scoreList[int(word[i] - 0x61)]
	}
	return score
}

func useLetters(word string, letterCounts map[byte]int) map[byte]int {
	updatedCounts := make(map[byte]int)
	for k,v := range letterCounts {
		updatedCounts[k] = v
	}
	for i, _ := range word {
		updatedCounts[word[i]]--
	}
	return updatedCounts
}

func countLetters(letters []byte) map[byte]int {
	letterCounts := make(map[byte]int)
	for _, v := range letters {
		letterCounts[v]++
	}
	return letterCounts
}

func optimalScore(words []string, letterCounts map[byte]int, score []int) int {
	if len(words) == 0 {
		return 0
	}
	var scoreIncludingWord int
	if canMakeWord(words[0], letterCounts) {
		scoreIncludingWord = computeWordScore(words[0], score) + optimalScore(words[1:], useLetters(words[0], letterCounts), score)
	}
	scoreExcludingWord := optimalScore(words[1:], letterCounts, score)
	if scoreIncludingWord > scoreExcludingWord {
		return scoreIncludingWord
	}
	return scoreExcludingWord
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	letterCounts := countLetters(letters)
	return optimalScore(words, letterCounts, score)
}
