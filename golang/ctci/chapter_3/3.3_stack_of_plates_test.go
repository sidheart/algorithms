package chapter_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetOfStacks(t *testing.T) {
	stack := NewSetOfStacks(2)
	assert.True(t, stack.IsEmpty())
	stack.Push(1)
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 1, len(stack.stacks))
	assert.Equal(t, 1, stack.Peek())
	stack.Push(2)
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 1, len(stack.stacks))
	assert.Equal(t, 2, stack.Peek())
	stack.Push(3)
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 2, len(stack.stacks))
	assert.Equal(t, 3, stack.Pop())
	assert.Equal(t, 1, len(stack.stacks))
}
