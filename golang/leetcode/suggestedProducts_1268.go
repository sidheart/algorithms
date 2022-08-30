package leetcode

import (
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	var answer [][]string
	sort.Sort(sort.StringSlice(products))
	prefix := ""
	for i := 0; i < len(searchWord); i++ {
		prefix += string(searchWord[i])
		var curSuggestion []string
		for _, product := range products {
			if strings.HasPrefix(product, prefix) && len(curSuggestion) < 3 {
				curSuggestion = append(curSuggestion, product)
			}
			// Can improve by checking whether we're lexicographically beyond the prefix here
			if len(curSuggestion) == 3 {
				break
			}
		}
		answer = append(answer, curSuggestion)
	}
	return answer
}
