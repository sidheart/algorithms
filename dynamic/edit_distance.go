package dynamic

import (
	"github.com/sidheart/algorithms/util"
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

type conversionCost func (a byte, b byte, action Action) int

func editDistance(x string, y string, xIndex int, yIndex int, costFunc conversionCost, memo map[util.Cell]solution) int {
	n := len(x) - xIndex
	m := len(y) - yIndex
	cell := util.Cell{xIndex, yIndex}
	if n <= 0 && m <= 0 {
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
	insertCost := solution{editDistance(x, y, xIndex, yIndex+1, costFunc, memo) + costFunc(a, b, INSERT), INSERT}
	deleteCost := solution{editDistance(x, y, xIndex+1, yIndex, costFunc, memo) + costFunc(a, b, DELETE), DELETE}
	replaceCost := solution{editDistance(x, y, xIndex+1, yIndex+1, costFunc, memo) + costFunc(a, b, REPLACE), REPLACE}
	memo[cell] = min(insertCost, deleteCost, replaceCost)
	return memo[cell].cost
}