package util

// Matrix is a rough representation of a 2 dimensional matrix, it is assumed that all rows are of uniform length
type Matrix [][]interface{}

// Cell represents the 0-indexed location of an element of a Matrix
// e.g. for an nxn Matrix A, Cell{1,2} represents the 2nd row and the 3rd column
type Cell struct {
	Row int
	Col int
}

// Returns the dimensions (rows by columns) of a Matrix
func (matrix *Matrix) Dimensions() (int, int) {
	derefMatrix := *matrix
	if derefMatrix == nil || len(derefMatrix) == 0 {
		return 0, 0
	}
	return len(derefMatrix), len(derefMatrix[0])
}