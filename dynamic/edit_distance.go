package dynamic

import (
	"github.com/sidheart/algorithms/util"
	"math"
)

type Action int

const (
	NONE Action = iota
	INSERT
	DELETE
	REPLACE
)

type solution struct {
	cost int
	action Action
}

func maxSolution() solution {
	return solution{math.MaxInt64, NONE}
}

func min(data  ...solution) solution {
	if len(data) == 0 {
		return solution{}
	}
	min := data[0]
	for _, num := range data {
		if num.cost < min.cost {
			min = num
		}
	}
	return min
}

// The cost of converting a character a into the character b by performing the specified action
type conversionCost func (a byte, b byte, action Action) int

func simpleCost(a byte, b byte, action Action) int {
	if a == b {
		return 0
	}
	switch action {
	case INSERT:
		return 1
	case DELETE:
		return 1
	case REPLACE:
		return 1
	case NONE:
		return math.MaxInt64
	default:
		return math.MaxInt64
	}
}

func EditDistance(x string, y string) int {
	solution := make(map[util.Cell]solution, len(x) * len(y))
	return editDistance(x, y, 0, 0, simpleCost, solution)
}

func editDistance(x string, y string, xIndex int, yIndex int, costFunc conversionCost, memo map[util.Cell]solution) int {
	if xIndex < 0 || yIndex < 0 {
		panic("Negative indices are not supported")
	}
	n := len(x)
	m := len(y)
	cell := util.Cell{xIndex, yIndex}
	sol, ok := memo[cell]
	if ok {
		return sol.cost
	}
	if n - xIndex <= 0 && m - yIndex <= 0 {
		memo[cell] = solution{0, NONE}
		return memo[cell].cost
	}
	var a, b byte
	if xIndex < len(x)  {
		a = x[xIndex]
	}
	if yIndex < len(y) {
		b = y[yIndex]
	}
	insertCost, deleteCost := maxSolution(), maxSolution()
	if b != 0 {
		insertCost = solution{editDistance(x, y, xIndex, yIndex+1, costFunc, memo) + costFunc(a, b, INSERT), INSERT}
	}
	if a != 0 {
		deleteCost = solution{editDistance(x, y, xIndex+1, yIndex, costFunc, memo) + costFunc(a, b, DELETE), DELETE}
	}
	replaceCost := solution{editDistance(x, y, xIndex+1, yIndex+1, costFunc, memo) + costFunc(a, b, REPLACE), REPLACE}
	memo[cell] = min(insertCost, deleteCost, replaceCost)
	return memo[cell].cost
}