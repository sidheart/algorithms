package chapter_3

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var eInt interface{}

func TestMultiStack_IsEmpty(t *testing.T) {
	threeStack := NewThreeInOneStack(10)
	e, _ := threeStack.IsEmpty(0)
	assert.True(t, e)
	e, _ = threeStack.IsEmpty(1)
	assert.True(t, e)
	e, _ = threeStack.IsEmpty(2)
	assert.True(t, e)
	threeStack.Push(1, 2)
	e, _ = threeStack.IsEmpty(1)
	assert.False(t, e)
}

func TestMultiStack_Push(t *testing.T) {
	stack := NewThreeInOneStack(5)
	stack.Push(0, 1)
	stack.Push(0, 2)
	stack.Push(0, 3)
	stack.Push(0, 4)
	stack.Push(0, 5)
	stack.Push(1, 1)
	stack.Push(1, 2)
	stack.Push(1, 3)
	stack.Push(1, 4)
	stack.Push(1, 5)
	stack.Push(2, 1)
	stack.Push(2, 2)
	stack.Push(2, 3)
	stack.Push(2, 4)
	stack.Push(2, 5)
	fmt.Println(stack.stack)
	expectedSegment := []interface{}{1, 2, 3, 4, 5}
	for i := 0; i < 15; i += 5 {
		assert.Equal(t, expectedSegment, stack.stack[i:i+5])
	}
}

func TestMultiStack_Grow(t *testing.T) {
	stack := NewThreeInOneStack(5)
	stack.Push(0, 1)
	stack.Push(0, 2)
	stack.Push(0, 3)
	stack.Push(0, 4)
	stack.Push(0, 5)
	stack.Push(1, 1)
	stack.Push(1, 2)
	stack.Push(1, 3)
	stack.Push(1, 4)
	stack.Push(1, 5)
	stack.Push(2, 1)
	stack.Push(2, 2)
	stack.Push(2, 3)
	stack.Push(2, 4)
	stack.Push(2, 5)
	stack.Push(1, 6)
	expectedSegment := []interface{}{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6, eInt, eInt, eInt, eInt, eInt, 1, 2, 3, 4, 5}
	assert.Equal(t, expectedSegment, stack.stack)
	stack.Push(2, 6)
	expectedSegment = []interface{}{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 6, eInt, eInt, eInt, eInt, eInt, 1, 2, 3, 4, 5, 6, eInt, eInt, eInt, eInt, eInt}
	assert.Equal(t, expectedSegment, stack.stack)
	stack.Push(0, 6)
	expectedSegment = []interface{}{1, 2, 3, 4, 5, 6, eInt, eInt, eInt, eInt, eInt, 1, 2, 3, 4, 5, 6, eInt, eInt, eInt, eInt, eInt, 1, 2, 3, 4, 5, 6, eInt, eInt, eInt, eInt, eInt}
	assert.Equal(t, expectedSegment, stack.stack)
}

func TestMultiStack_Pop(t *testing.T) {
	stack := NewThreeInOneStack(1)
	stack.Push(0, 1)
	stack.Push(0, 2)
	stack.Push(1, 1)
	stack.Push(1, 2)
	stack.Push(2, 1)
	stack.Push(2, 2)
	stack.Pop(0)
	stack.Pop(1)
	stack.Pop(2)
	var x interface{}
	x, _ = stack.Peek(0)
	assert.Equal(t, 1, x)
	x, _ = stack.Peek(1)
	assert.Equal(t, 1, x)
	x, _ = stack.Peek(2)
	assert.Equal(t, 1, x)
	stack.Push(0, 3)
	stack.Push(1, 3)
	stack.Push(2, 3)
	x, _ = stack.Peek(0)
	assert.Equal(t, 3, x)
	x, _ = stack.Peek(1)
	assert.Equal(t, 3, x)
	x, _ = stack.Peek(2)
	assert.Equal(t, 3, x)
}

func TestMultiStack_Peek(t *testing.T) {
	stack := NewThreeInOneStack(1)
	top, err := stack.Peek(0)
	assert.NotNil(t, err)
	assert.Nil(t, top)
	assert.Equal(t, "attempted to peek on empty stack", err.Error())
	stack.Push(0, 1)
	stack.Push(1, 2)
	stack.Push(2, 3)
	var x interface{}
	x, _ = stack.Peek(0)
	assert.Equal(t, 1, x)
	x, _ = stack.Peek(1)
	assert.Equal(t, 2, x)
	x, _ = stack.Peek(2)
	assert.Equal(t, 3, x)
	stack.Push(0, 4)
	stack.Push(1, 5)
	stack.Push(2, 6)
	x, _ = stack.Peek(0)
	assert.Equal(t, 4, x)
	x, _ = stack.Peek(1)
	assert.Equal(t, 5, x)
	x, _ = stack.Peek(2)
	assert.Equal(t, 6, x)
	stack.Pop(0)
	stack.Pop(1)
	stack.Pop(2)
	x, _ = stack.Peek(0)
	assert.Equal(t, 1, x)
	x, _ = stack.Peek(1)
	assert.Equal(t, 2, x)
	x, _ = stack.Peek(2)
	assert.Equal(t, 3, x)
}
