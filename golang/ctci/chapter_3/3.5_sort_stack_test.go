package chapter_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortStack(t *testing.T) {
	stack := NewSortStack(5)
	assert.True(t, stack.IsEmpty())
	stack.Push(3)
	assert.False(t, stack.IsEmpty())
	stack.Push(1)
	stack.Push(4)
	stack.Push(5)
	stack.Push(2)
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, stack.Peek())
		assert.Equal(t, i, stack.Pop())
	}
	assert.True(t, stack.IsEmpty())
}
