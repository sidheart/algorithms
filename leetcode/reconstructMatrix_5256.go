package leetcode

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	matrix := make([][]int, 2)
	m := len(colsum)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	numTwos := 0
	numOnes := 0
	for _, sum := range colsum {
		if sum == 2 {
			numTwos++
		} else if sum == 1 {
			numOnes++
		}
	}
	remainingUpper, remainingLower := upper - numTwos, lower - numTwos
	if remainingUpper < 0 || remainingLower < 0 || numOnes != remainingUpper + remainingLower {
		return [][]int{}
	}
	for col, sum := range colsum {
		if sum == 2 {
			matrix[0][col] = 1
			matrix[1][col] = 1
		} else if sum == 1 {
			if remainingUpper > 0 {
				matrix[0][col] = 1
				remainingUpper--
			} else if remainingLower > 0 {
				matrix[1][col] = 1
				remainingLower--
			} else {
				return [][]int{}
			}
		}
	}
	return matrix
}
