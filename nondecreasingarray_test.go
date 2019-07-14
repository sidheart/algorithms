package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCheckPossibility(t *testing.T) {
	assert.True(t, checkPossibility([]int{1, 2, 3, 4}), "basic truth")
	assert.False(t, checkPossibility([]int{4, 3, 2, 1}), "basic falseness")
	assert.True(t, checkPossibility([]int{1, 2, 3, 4, 5}), "odd truth")
	assert.False(t, checkPossibility([]int{5, 4, 3, 2, 1}), "odd falseness")

	assert.True(t, checkPossibility([]int{2, 2, 2}), "less than or equals I")
	assert.False(t, checkPossibility([]int{2, 1, 1, 0}), "less than or equals III")
	assert.True(t, checkPossibility([]int{2, 1, 1 ,2}), "less than or equals II")

	assert.True(t, checkPossibility([]int{2, 1, 1 ,2}), "change beginning")
	assert.True(t, checkPossibility([]int{1, 2, 2 ,1}), "change end")
	assert.True(t, checkPossibility([]int{2, 1, 1, 1 ,2}), "odd change beginning")
	assert.True(t, checkPossibility([]int{1, 2, 2, 2 ,1}), "odd change end")

	assert.False(t, checkPossibility([]int{3, 4, 2, 3}), "oscillation I")
	assert.False(t, checkPossibility([]int{3, 2, 4, 3}), "oscillation II")
	assert.True(t, checkPossibility([]int{3, 4, 2, 4}), "oscillation III")




}
