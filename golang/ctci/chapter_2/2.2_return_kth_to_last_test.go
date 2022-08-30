package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnKthToLast(t *testing.T) {
	head := MakeLinkedList([]interface{}{0, 1, 2})
	assert.Equal(t, 0, returnKthToLast(head, 2).Value)
	assert.Equal(t, 1, returnKthToLast(head, 1 ).Value)
	assert.Equal(t, 2, returnKthToLast(head, 0).Value)
	var empty *ListNode
	assert.Equal(t, empty, returnKthToLast(head, 3))
	head = nil
	assert.Equal(t, empty, returnKthToLast(head, 0))
}
