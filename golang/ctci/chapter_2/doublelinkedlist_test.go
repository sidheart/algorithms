package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var empty *DoubleListNode

func TestMakeDoublyLinkedList(t *testing.T) {
	head := MakeDoublyLinkedList([]interface{}{1, 2, 3, 4, 5})
	runner := head
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, runner.Value)
		runner = runner.Next
	}
	assert.Equal(t, empty, runner)
	tail := head
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}
	for i := 5; i >= 1; i-- {
		assert.Equal(t, i, tail.Value)
		tail = tail.Prev
	}
	assert.Equal(t, empty, tail)
}
