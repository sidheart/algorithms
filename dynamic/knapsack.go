package dynamic

import (
	"github.com/sidheart/algorithms/util"
	"math"
)

type item struct {
	weight, value int
}

// optimalKnapsack represents a full knapsack with a total value and remaining weight
type optimalKnapsack item

func Knapsack(items []item, maxWeight int) int {
	memo := make(map[util.Cell]optimalKnapsack)
	return knapsack(items, len(items), maxWeight, 0, memo)
}

func knapsack(items []item, n int, maxWeight int, i int, memo map[util.Cell]optimalKnapsack) int {
	if i >= n {
		return 0
	}
	curCell := util.Cell{Row: i, Col: maxWeight}
	curItem := items[i]
	if n - i == 1 {
		if curItem.weight <= maxWeight {
			memo[curCell] = optimalKnapsack{curItem.value, maxWeight - curItem.weight}
		} else {
			memo[curCell] = optimalKnapsack{0, maxWeight}
		}
	}
	solution, ok := memo[curCell]
	if ok {
		return solution.value
	}
	includingItem := math.MinInt64
	if curItem.weight <= maxWeight {
		includingItem = curItem.value + knapsack(items, n, maxWeight-curItem.weight, i+1, memo)
	}
	notIncludingItem := knapsack(items, n, maxWeight, i+1, memo)
	if includingItem > notIncludingItem {
		memo[curCell] = optimalKnapsack{maxWeight - curItem.weight, includingItem}
		return includingItem
	}
	memo[curCell] = optimalKnapsack{maxWeight, notIncludingItem}
	return notIncludingItem
}
