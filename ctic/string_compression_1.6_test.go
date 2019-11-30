package ctic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompress(t *testing.T) {
	s := "aaabcccccaaa"
	assert.Equal(t, "a3b1c5a3", compress(s))
	s = "aaabccccca"
	assert.Equal(t, "a3b1c5a1", compress(s))
	s = "abc"
	assert.Equal(t, "abc", compress(s))
	s = ""
	assert.Equal(t, "", compress(s))
	s = "aaaaaaaaaabbbbbbbbbbbbbbb"
	assert.Equal(t, "a10b15", compress(s))
}
