package ctci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testMatrixA = [][]int{
	{1, 0, 1},
	{0, 1, 1},
	{1, 1, 1},
}

var resMatrixA = [][]int{
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 1},
}

var testMatrixB = [][]int{
	{1, 1, 1},
	{1, 0, 1},
	{1, 1, 1},
}

var resMatrixB = [][]int{
	{1, 0, 1},
	{0, 0, 0},
	{1, 0, 1},
}

var testMatrixC = [][]int{
	{1, 1, 0},
}

var resMatrixC = [][]int{
	{0, 0, 0},
}

var testMatrixD = [][]int{
	{1, 1, 1},
	{1, 1, 1},
}

var resMatrixD = testMatrixD

var testMatrixE = [][]int{
	{1, 1, 1},
	{1, 1, 1},
	{1, 1, 0},
}

var resMatrixE = [][]int{
	{1, 1, 0},
	{1, 1, 0},
	{0, 0, 0},
}

func TestZeroMatrix(t *testing.T) {
	assert.Equal(t, resMatrixA, zeroMatrix(testMatrixA))
	assert.Equal(t, resMatrixB, zeroMatrix(testMatrixB))
	assert.Equal(t, resMatrixC, zeroMatrix(testMatrixC))
	assert.Equal(t, resMatrixD, zeroMatrix(testMatrixD))
	assert.Equal(t, resMatrixE, zeroMatrix(testMatrixE))
}