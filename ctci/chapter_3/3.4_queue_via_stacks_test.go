package chapter_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueViaStacks(t *testing.T) {
	q := NewStackQueue(5)
	assert.True(t, q.IsEmpty())
	q.Add(1)
	q.Add(2)
	q.Add(3)
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 1, q.Remove())
	q.Add(4)
	assert.Equal(t, 2, q.Peek())
	assert.Equal(t, 2, q.Remove())
	assert.Equal(t, 3, q.Remove())
	assert.Equal(t, 4, q.Remove())
	assert.True(t, q.IsEmpty())
}
