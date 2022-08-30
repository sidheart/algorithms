package leetcode

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	var start, longest int
	for end := 0; end < len(s); end++ {
		index, repeat := charMap[s[end]]
		if repeat  && index > start {
			start = index
		}
		charMap[s[end]] = end + 1
		curLen := (end - start) + 1
		if curLen > longest {
			longest = curLen
		}
	}
	return longest
}
