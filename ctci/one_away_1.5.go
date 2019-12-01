package ctci

import "math"

// countDeletions returns the number of deletions necessary to transform s1 into s2
// if transformation is impossible simply through deletion from s1, this algorithm will return the number of attempted deletions
// s1 must be a string with greater length than s2
func countDeletions(s1, s2 string) (differences int) {
	m, n := len(s1), len(s2)
	if n == 0 {
		return m
	}
	i, j := 0, 0
	for i < m && j < n {
		if s1[i] != s2[j] {
			differences++
			i++
		} else {
			i++
			j++
		}
	}
	differences += m - i
	differences += n - j
	return differences
}

// isOneAway returns true if s1 can be transformed into s2 with only a single insert, delete, or replace operation
// This function runs in O(M+N) time
func isOneAway(s1, s2 string) bool {
	m, n := len(s1), len(s2)
	if m == n {
		differences := 0
		for i := 0; i < m; i++ {
			if s1[i] != s2[i] {
				differences++
			}
			if differences > 1 {
				return false
			}
		}
		return true
	}
	if math.Abs(float64(m - n)) == 1 {
		if m > n {
			return countDeletions(s1, s2) <= 1
		}
		return countDeletions(s2, s1) <= 1
	}
	return false
}
