package chapter_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUnique(t *testing.T) {
	assert.True(t, isUnique("abcAde大岛地"))
	assert.False(t, isUnique("abca"))
	assert.True(t, isUniqueAscii("AaBbCcDd"))
	assert.False(t, isUniqueAscii("ABbAcde"))
}
