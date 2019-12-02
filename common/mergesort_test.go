package common

import (
	"github.com/sidheart/algorithms/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// Check that MergeSort works when it is passed the empty list
	assert.Equal(t, []util.Comparable{}, MergeSort([]util.Comparable{}))
	// Check that MergeSort works on single element lists
	v := []util.Comparable{util.CmpInt(1)}
	assert.Equal(t, v, MergeSort(v))
	// Check that MergeSort works on lists of odd length
	sorted := []util.Comparable{util.CmpInt(1), util.CmpInt(2), util.CmpInt(3)}
	v = []util.Comparable{util.CmpInt(3), util.CmpInt(2), util.CmpInt(1)}
	assert.Equal(t, sorted, MergeSort(v))
	assert.NotEqual(t, sorted, v)
	v = []util.Comparable{util.CmpInt(3), util.CmpInt(1), util.CmpInt(2)}
	assert.Equal(t, sorted, MergeSort(v))
	// Check that MergeSort works on lists of even length
	sorted = []util.Comparable{util.CmpInt(1), util.CmpInt(2), util.CmpInt(3), util.CmpInt(4)}
	v = []util.Comparable{util.CmpInt(4), util.CmpInt(3), util.CmpInt(2), util.CmpInt(1)}
	assert.Equal(t, sorted, MergeSort(v))
	v = []util.Comparable{util.CmpInt(1), util.CmpInt(4), util.CmpInt(2), util.CmpInt(3)}
	assert.Equal(t, sorted, MergeSort(v))
	// Check that MergeSort works with lists containing duplicate elements
	sorted = []util.Comparable{util.CmpInt(1), util.CmpInt(1), util.CmpInt(3), util.CmpInt(4)}
	v = []util.Comparable{util.CmpInt(3), util.CmpInt(1), util.CmpInt(4), util.CmpInt(1)}
	assert.Equal(t, sorted, MergeSort(v))
	sorted = []util.Comparable{util.CmpInt(1), util.CmpInt(1), util.CmpInt(3)}
	v = []util.Comparable{util.CmpInt(1), util.CmpInt(3), util.CmpInt(1)}
	assert.Equal(t, sorted, MergeSort(v))
}
