package chapter_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPalindrome(t *testing.T) {
	valid1 := MakeLinkedList([]interface{}{1, 0, 1})
	valid2 := MakeLinkedList([]interface{}{0, 0})
	invalid1 := MakeLinkedList([]interface{}{0, 0, 7})
	invalid2 := MakeLinkedList([]interface{}{1, 0})
	invalid3 := MakeLinkedList([]interface{}{1, 0, 1, 1})
	assert.True(t, palindrome(valid1))
	assert.True(t, palindrome(valid2))
	assert.False(t, palindrome(invalid1))
	assert.False(t, palindrome(invalid2))
	assert.False(t, palindrome(invalid3))
}
