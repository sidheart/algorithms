package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectLoop(t *testing.T) {
	c := MakeLinkedList([]interface{}{1, 2, 3, 4, 5})
	loop := c.Next.Next
	cTail := c
	for cTail.Next != nil {
		cTail = cTail.Next
	}
	cTail.Next = loop
	assert.Equal(t, loop, detectLoop(c))
}
