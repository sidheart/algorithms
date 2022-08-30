// Package dynamic implements various algorithms using Dynamic Programming.
package dynamic

// KnapsackItem represents an item that one might take with them in their knapsack
type KnapsackItem interface {
	Value() uint32
	Weight() uint32
}

// Knapsack computes the indices of a subset of items with highest total value that doesn't exceed a max Weight
func Knapsack(items []KnapsackItem, maxWeight uint32) (bestItems []uint32) {
	if items == nil || maxWeight < 0 {
		return bestItems
	}
	n := len(items)
	if n == 0 {
		return bestItems
	}
	memoTable := make([]uint32, n)
	includedItems := make([]bool, n)
	knapsack(0, maxWeight, uint32(n), &items, &memoTable, &includedItems)
	for i := 0; i < len(includedItems); i++ {
		if includedItems[i] {
			bestItems = append(bestItems, uint32(i))
		}
	}
	return
}

// knapsack computes the total value of the "best" subset of items with indices >= i
// "best" refers to the subset of items with largest total value that doesn't exceed w
func knapsack(i uint32, w uint32, n uint32, items *[]KnapsackItem, MemoTable *[]uint32, IncludedItems *[]bool) uint32 {
	if i >= n || w <= 0 {
		return 0
	}
	currentItem := (*items)[i]
	currentWeight := currentItem.Weight()
	var valueWithItem uint32
	if currentWeight<= w {
		valueWithItem = currentItem.Value() + knapsack(i + 1, w - currentWeight, n, items, MemoTable, IncludedItems)
	}
	valueWithoutItem := knapsack(i + 1, w, n, items, MemoTable, IncludedItems)
	if valueWithItem > valueWithoutItem {
		(*IncludedItems)[i] = true
		return valueWithItem
	}
	return valueWithoutItem
}
