package dynamic

// A Justifier computes the best way to split text into lines
// badnessFunc is a function that computes the justification "badness" of a line
// lineSplits is an array of indices representing where each line should end.
// e.g. if lineSplits[0] is 12, then the first line of text should consist of
// words 0 to 11 (inclusive) and so on.
type Justifier func(text []string, badnessFunc func([]string) (badness int)) (lineSplits []int)

// TextJustifier is a closure that returns a Justifier which can be used to justify text
func TextJustifier() Justifier {
	var memo []int
	var parent []int
	return func(text []string, badnessFunc func([]string) (badness int)) (lineSplits []int) {
		n := len(text)
		memo = make([]int, n+1)
		parent = make([]int, n+1)
		memo[n] = 0
		for i := n - 1; i >= 0; i-- {
			min := 0x7FFFFFFF // Maximum 32 bit signed integer
			for j := i + 1; j <= n; j++ {
				curMin := badnessFunc(text[i:j]) + memo[j]
				if curMin < min {
					min = curMin
					parent[i] = j
				}
			}
			memo[i] = min
		}
		for i := 0; i < n+1; {
			lineSplits = append(lineSplits, parent[i])
			i = parent[i]
		}
		return
	}
}
