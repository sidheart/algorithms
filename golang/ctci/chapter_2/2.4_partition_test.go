package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func verifyPartition(head *ListNode, k int) bool {
	left := true
	for head != nil {
		if head.Value.(int) < k && !left {
			return false
		}
		if head.Value.(int) >= k {
			left = false
		}
		head = head.Next
	}
	return true
}

func TestVerifyPartition(t *testing.T) {
	valid1 := MakeLinkedList([]interface{}{1, 3, 3, 4, 5})
	valid2 := MakeLinkedList([]interface{}{2, 1, 5, 3, 3})
	valid3 := MakeLinkedList([]interface{}{3, 3, 3, 3, 3})
	var valid4 *ListNode
	invalid1 := MakeLinkedList([]interface{}{1, 3, 3, 2, 5})
	assert.True(t, verifyPartition(valid1, 3))
	assert.True(t, verifyPartition(valid2, 3))
	assert.True(t, verifyPartition(valid3, 3))
	assert.True(t, verifyPartition(valid4, 3))
	assert.True(t, verifyPartition(invalid1, 7))
	assert.False(t, verifyPartition(invalid1, 3))
}

func TestPartition(t *testing.T) {
	head := MakeLinkedList([]interface{}{5, 4, 3, 2, 1})
	assert.True(t, verifyPartition(partition(head, 3), 3))
}
