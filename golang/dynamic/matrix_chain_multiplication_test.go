package dynamic

import (
	"github.com/sidheart/algorithms/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createZeroMatrix(rows int, cols int) util.Matrix {
	matrix := make(util.Matrix, rows)
	for i := 0; i < rows; i++ {
		s := make([]int, cols)
		matrix[i] = make([]interface{}, cols)
		for j, _ := range s {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func TestParenthesize(t *testing.T) {
	A := createZeroMatrix(4, 3)
	result := Parenthesize(&[]util.Matrix{A})
	assert.Equal(t, 0, result)

	B := createZeroMatrix(3, 2)
	result = Parenthesize(&[]util.Matrix{A, B})
	assert.Equal(t, 24, result)

	C := createZeroMatrix(2, 1)
	result = Parenthesize(&[]util.Matrix{A, B, C})
	assert.Equal(t, 18, result)
}

