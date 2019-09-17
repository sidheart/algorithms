package dynamic

import (
	"fmt"
	"github.com/sidheart/algorithms/util"
	"math"
)

// A multiplicationResult represents the cost of multiplying 2 matrices and the index between the two matrices
// e.g. for a list of Matrices [A_0, A_1] the cost would be cost(A_0, A_1) and the split would be 1
type multiplicationResult struct {
	cost int
	split int
}

// A costFunction returns the cost of multiplying Matrices A[i~k] and A[k~j]
// A[i~j] denotes the Matrix obtained by multiplying all the Matrices from A_i to A_j-1
type costFunction func (terms *[]util.Matrix, i int, j int, k int) int


// multiplicationCost is a costFunction that computes the cost of multiplying A[i~k] * A[k~j]
// Note that the cost of computing the individual terms A[i~k] and A[k~j] is not included
func mutliplicationCost(terms *[]util.Matrix, i int, j int, k int) int  {
	derefTerms := *terms
	rowLHS, _ := derefTerms[i].Dimensions()
	_, colLHS := derefTerms[k-1].Dimensions()
	_, colRHS := derefTerms[j].Dimensions()
	return rowLHS * colLHS * colRHS
}

// Returns true IFF it is possible to multiply every Matrix in terms to compute a single product Matrix
func mutliplicationPossible(terms *[]util.Matrix) bool {
	derefTerms := *terms
	for i := 1; i < len(derefTerms); i++ {
		_, colLHS := derefTerms[i-1].Dimensions()
		rowRHS, _ := derefTerms[i].Dimensions()
		if colLHS != rowRHS {
			return false
		}
	}
	return true
}

// Returns a parenthesized string representing the optimal way to multiply the given Matrices in terms
func printOptimalMultiplication(lo int, hi int, solution *map[util.Cell]multiplicationResult, terms *[]util.Matrix) string {
	if hi <= lo {
		return ""
	} else if hi - lo == 1 {
		return fmt.Sprintf("%v", (*terms)[lo])
	}
	k := (*solution)[util.Cell{lo, hi}].split
	return fmt.Sprintf("(%v * %v)", printOptimalMultiplication(lo, k, solution, terms), printOptimalMultiplication(k, hi, solution, terms))
}

// Computes the optimal way to compute A[i~j] and stores the result in memo[Cell{i,j}]
func parenthesize(terms *[]util.Matrix, i int, j int, costFunc costFunction, memo map[util.Cell]multiplicationResult) multiplicationResult {
	min := math.MaxInt64
	desiredCell := util.Cell{i,j}
	length := j - i
	if length < 0 {
		panic("j is greater than i for recursive call of parenthesize")
	} else if length <= 1 {
		memo[desiredCell] = multiplicationResult{0, -1}
		return memo[desiredCell]
	}
	for k := i + 1; k < j; k++ {
		optimalLHS, ok := memo[util.Cell{i, k}]
		if !ok {
			optimalLHS = parenthesize(terms, i, k, costFunc, memo)
		}
		optimalRHS, ok := memo[util.Cell{k, j - 1}]
		if !ok {
			optimalRHS = parenthesize(terms, k, j, costFunc, memo)
		}
		potentialMin :=  optimalLHS.cost + optimalRHS.cost + costFunc(terms, i, j-1, k)
		if potentialMin < min {
			min = potentialMin
			memo[desiredCell] = multiplicationResult{min, k}
		}
	}
	return memo[desiredCell]
}

// Returns the minimum number of multiplications required to multiply every Matrix in terms and prints an expression
// that achieves this minimum number when computed
func Parenthesize(terms *[]util.Matrix) int {
	n := len(*terms)
	memo := make(map[util.Cell]multiplicationResult, n * n)
	if !mutliplicationPossible(terms) {
		panic("Matrix multiplication is not possible for the provided arguments")
	}
	solution := parenthesize(terms, 0, n, mutliplicationCost, memo)
	fmt.Println(printOptimalMultiplication(0, n, &memo, terms))
	return solution.cost
}
