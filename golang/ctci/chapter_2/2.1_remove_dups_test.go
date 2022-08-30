package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	head := MakeLinkedList([]interface{}{1, 1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5})
	expected := MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	removeDups(head)
	assert.True(t, head.Equals(expected))
	head = nil
	expected = nil
	removeDups(head)
	assert.True(t, head.Equals(expected))
	head = MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	expected = MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	removeDups(head)
	assert.True(t, head.Equals(expected))
	head = MakeLinkedList([]interface{}{1, 1})
	expected = MakeLinkedList([]interface{}{1})
	removeDups(head)
	assert.True(t, head.Equals(expected))
}

func TestRemoveDupsNoBuffer(t *testing.T) {
	head := MakeLinkedList([]interface{}{1, 1, 2, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5})
	expected := MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	removeDupsNoBuffer(head)
	assert.True(t, head.Equals(expected))
	head = nil
	expected = nil
	removeDupsNoBuffer(head)
	assert.True(t, head.Equals(expected))
	head = MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	expected = MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	removeDupsNoBuffer(head)
	assert.True(t, head.Equals(expected))
	head = MakeLinkedList([]interface{}{1, 1})
	expected = MakeLinkedList([]interface{}{1})
	removeDupsNoBuffer(head)
	assert.True(t, head.Equals(expected))
}
