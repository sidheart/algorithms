package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteMiddleNode(t *testing.T) {
	head := MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	expected := MakeLinkedList([]interface{}{1, 2, 4, 5})
	deleteMiddleNode(head.Next.Next)
	assert.Equal(t, expected, head)
	expected = MakeLinkedList([]interface{}{1, 4, 5})
	deleteMiddleNode(head.Next)
	assert.Equal(t, expected, head)
}
