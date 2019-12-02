package common

import (
	"github.com/sidheart/algorithms/util"
)

func merge(a, b []util.Comparable) (c []util.Comparable) {
	m, n := len(a), len(b)
	if m == 0 {
		copy(c, b)
		return c
	} else if n == 0 {
		copy(c, a)
		return c
	}
	var i, j int
	for i < m && j < n {
		if a[i].Less(b[j]) {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	if i < m {
		c = append(c, a[i:]...)
	} else if j < n {
		c = append(c, b[j:]...)
	}
	return c
}

func MergeSort(a []util.Comparable) []util.Comparable {
	n := len(a)
	if n <= 1 {
		return a
	}
	mid := n / 2
	return merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}
