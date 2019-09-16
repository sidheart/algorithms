package dynamic

import "math"

// Matrix is a rough representation of a 2 dimensional matrix, it is assumed that all rows are of uniform length
type Matrix [][]interface{}

type tuple struct {
	x int
	y int
}

type multiplicationResult struct {
	cost int
	split int
}

// A costFunction returns the cost of multiplying Matrices A[i~k] and A[k~j]
// A[i~j] denotes the Matrix obtained by multiplying all the Matrices from terms[i] to terms[j-1]
type costFunction func (terms *[]Matrix, i int, j int, k int) int

// Returns the dimensions (rows by columns) of a Matrix
func (matrix *Matrix) dimensions() (int, int) {
	derefMatrix := *matrix
	if derefMatrix == nil || len(derefMatrix) == 0 {
		return 0, 0
	}
	return len(derefMatrix), len(derefMatrix[0])
}

func mutliplicationCost(terms *[]Matrix, i int, j int, k int) int  {
	derefTerms := *terms
	rowLHS, _ := derefTerms[i].dimensions()
	_, colLHS := derefTerms[k].dimensions()
	_, colRHS := derefTerms[j].dimensions()
	return rowLHS * colLHS * colRHS
}

func mutliplicationPossible(terms *[]Matrix) bool {
	derefTerms := *terms
	for i := 1; i < len(derefTerms); i++ {
		_, colLHS := derefTerms[i-1].dimensions()
		rowRHS, _ := derefTerms[i].dimensions()
		if colLHS != rowRHS {
			return false
		}
	}
	return true
}

func parenthesize(terms *[]Matrix, i int, j int, costFunc costFunction, memo map[tuple]multiplicationResult) multiplicationResult {
	min := math.MinInt64
	cell := tuple{i,j}
	length := j - i
	if length < 0 {
		panic("j is greater than i for recursive call of parenthesize")
	} else if length <= 1 {
		memo[cell] = multiplicationResult{0, -1}
		return memo[cell]
	}
	for k := i + 1; k < j; k++ {
		optimalLHS, ok := memo[tuple{i, k}]
		if !ok {
			optimalLHS = parenthesize(terms, i, k, costFunc, memo)
		}
		optimalRHS, ok := memo[tuple{k, j - 1}]
		if !ok {
			optimalRHS = parenthesize(terms, k, j, costFunc, memo)
		}
		potentialMin :=  optimalLHS.cost + optimalRHS.cost + costFunc(terms, i, j, k)
		if potentialMin < min {
			min = potentialMin
			memo[cell] = multiplicationResult{min, k}
		}
	}
	return memo[cell]
}

func Parenthesize(terms *[]Matrix) int {
	derefTerms := *terms
	n := len(derefTerms)
	memo := make(map[tuple]multiplicationResult, n * n)
	if !mutliplicationPossible(terms) {
		panic("Matrix multiplication is not possible for the provided arguments")
	}
	return parenthesize(terms, 0, n, mutliplicationCost, memo).cost
}
