package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumLists(t *testing.T) {
	a := MakeLinkedList([]interface{}{7, 1, 6})
	b := MakeLinkedList([]interface{}{5, 9, 2})
	sum := MakeLinkedList([]interface{}{2, 1, 9})
	assert.True(t, sumLists(a, b).Equals(sum))
	a = MakeLinkedList([]interface{}{9, 9, 9})
	b = MakeLinkedList([]interface{}{9, 9})
	sum = MakeLinkedList([]interface{}{8,9,0,1})
	assert.True(t, sumLists(a, b).Equals(sum))
}

func TestSumListsForward(t *testing.T) {
	a := MakeLinkedList([]interface{}{7, 1, 6})
	b := MakeLinkedList([]interface{}{5, 9, 2})
	sum := MakeLinkedList([]interface{}{1, 3, 0, 8})
	assert.True(t, sumListsForward(a, b).Equals(sum))
	a = MakeLinkedList([]interface{}{9, 9, 9})
	b = MakeLinkedList([]interface{}{9, 9})
	sum = MakeLinkedList([]interface{}{1,0,9,8})
	assert.True(t, sumListsForward(a, b).Equals(sum))
}
