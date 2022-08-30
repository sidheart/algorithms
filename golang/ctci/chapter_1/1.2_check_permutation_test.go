package chapter_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	assert.True(t, CheckPermutation("abcd", "dbca"))
	assert.True(t, CheckPermutation("ab八cd", "dbc八a"))
	assert.False(t, CheckPermutation("八a", "a"))
	assert.False(t, CheckPermutation("a", "aA"))
	assert.False(t, CheckPermutation("a", "b"))
}
