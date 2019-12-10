package leetcode

func bestPalindrome(s string, center int) (bestLeft int, bestRight int) {
	var left, right int
	if center % 2 == 0 {
		left = center/2 - 1
		right = center/2
	} else {
		left = center/2 - 1
		right = center/2 + 1
	}
	for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
		if s[left] != s[right] {
			return bestLeft, bestRight
		}
		bestLeft, bestRight = left, right
	}
	return bestLeft, bestRight
}

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	bestLeft, bestRight := 0,0
	for i := 1; i <= 2 * n - 1; i++ {
		tempLeft, tempRight := bestPalindrome(s, i)
		if tempRight - tempLeft > bestRight - bestLeft {
			bestLeft, bestRight = tempLeft, tempRight
		}
	}
	return s[bestLeft:bestRight+1]
}
