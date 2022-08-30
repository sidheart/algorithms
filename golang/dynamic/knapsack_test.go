package dynamic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestItem struct {
	V, W uint32
}

func (t TestItem) Value() uint32 {
	return t.V
}

func (t TestItem) Weight() uint32 {
	return t.W
}

func TestKnapsack(t *testing.T) {
	assert.Equal(t, []uint32(nil), Knapsack(nil, 1000))
	assert.Equal(t, []uint32(nil), Knapsack([]KnapsackItem{TestItem{1, 2}}, 0))
	knapsack := []KnapsackItem{TestItem{2, 1}, 
		TestItem{1000, 99}, TestItem{2, 1}, TestItem{1000, 98}}
	assert.Equal(t, []uint32{0, 2, 3}, Knapsack(knapsack, 100))
}
