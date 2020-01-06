package chapter_1

// coord represents a coordinate in a given matrix
type coord struct {
	row, col int
}

// zeroCross sets every element of the cross centered at (row, col) to 0
func zeroCross(matrix [][]int, row, col int) {
	m, n := len(matrix), len(matrix[0])
	for curCol := 0; curCol < n; curCol++ {
		matrix[row][curCol] = 0
	}
	for curRow := 0; curRow < m; curRow++ {
		matrix[curRow][col] = 0
	}
}

// copyMatrix allocates a new matrix that is identical to m, and returns it
func copyMatrix(m [][]int) [][]int {
	var matrix [][]int
	for _, row := range m {
		newRow := make([]int, len(row))
		copy(newRow, row)
		matrix = append(matrix, newRow)
	}
	return matrix
}

// zeroMatrix returns a matrix such that for every 0 in the original matrix, the entire row and column of that 0 is set
// to 0 in the new matrix
func zeroMatrix(m [][]int) [][]int {
	if m == nil || len(m) == 0 || len(m[0]) == 0 {
		return m
	}
	matrix := copyMatrix(m)
	var coordsToZero []coord
	for row, _ := range matrix {
		for col, _ := range matrix[0] {
			if matrix[row][col] == 0 {
				coordsToZero = append(coordsToZero, coord{row, col})
			}
		}
	}
	for _, coord := range coordsToZero {
		zeroCross(matrix, coord.row, coord.col)
	}
	return matrix
}
