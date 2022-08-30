package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	s := MakeLinkedList([]interface{}{0, 1, 0, 1})
	a := MakeLinkedList([]interface{}{3, 7, 9})
	b := MakeLinkedList([]interface{}{1})
	assert.False(t, intersection(a, b))
	aItr := a
	for aItr.Next != nil {
		aItr = aItr.Next
	}
	aItr.Next = s
	b.Next = s
	assert.True(t, intersection(a, b))
}

func TestGetIntersection(t *testing.T) {
	s := MakeLinkedList([]interface{}{0, 1, 0, 1})
	a := MakeLinkedList([]interface{}{3, 7, 9})
	b := MakeLinkedList([]interface{}{1})
	var empty *ListNode
	assert.Equal(t, empty, getIntersection(a, b))
	aItr := a
	for aItr.Next != nil {
		aItr = aItr.Next
	}
	aItr.Next = s
	b.Next = s
	assert.Equal(t, s, getIntersection(a, b))
}
