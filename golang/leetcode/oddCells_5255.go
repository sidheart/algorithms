package leetcode

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, m)
	}
	oddCount := 0
	for _, index := range indices {
		rowToIncrement := index[0]
		colToIncrement := index[1]
		// Increment row
		for i := 0; i < len(matrix[0]); i++ {
			matrix[rowToIncrement][i] += 1
			if matrix[rowToIncrement][i] % 2 == 0 {
				oddCount -= 1
			} else {
				oddCount += 1
			}
		}
		// Increment col
		for i := 0; i < len(matrix); i++ {
			matrix[i][colToIncrement] += 1
			if matrix[i][colToIncrement] % 2 == 0 {
				oddCount -= 1
			} else {
				oddCount += 1
			}
		}
	}
	return oddCount
}
