package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeLinkedList(t *testing.T) {
	head := MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, head.Value)
		head = head.Next
	}
}

func TestListNode_Equals(t *testing.T) {
	head := MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	head2 := MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	assert.True(t, head.Equals(head2))
	assert.True(t, head2.Equals(head))
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, head.Value)
		head = head.Next
	}
	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, head2.Value)
		head2 = head2.Next
	}
	head = nil
	head2 = nil
	assert.True(t, head.Equals(head2))
	assert.True(t, head2.Equals(head))
	head = nil
	head2 = MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	assert.False(t, head.Equals(head2))
	assert.False(t, head2.Equals(head))
	head = MakeLinkedList([]interface{}{2, 2, 3, 4})
	head2 = MakeLinkedList([]interface{}{2, 2, 3, 4, 5})
	assert.False(t, head.Equals(head2))
	assert.False(t, head2.Equals(head))
}

func TestListNode_Len(t *testing.T) {
	head := MakeLinkedList([]interface{}{1, 1, 3, 4, 5})
	assert.Equal(t, 5, head.Len())
	head = nil
	assert.Equal(t, 0, head.Len())
}
