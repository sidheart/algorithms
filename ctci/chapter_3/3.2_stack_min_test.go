package chapter_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	stack := NewMinStack(10)
	stack.Push(1)
	assert.Equal(t, 1, stack.Min())
	stack.Push(3)
	assert.Equal(t, 1, stack.Min())
	stack.Push(-1)
	assert.Equal(t, -1, stack.Min())
	stack.Pop()
	assert.Equal(t, 1, stack.Min())
}
