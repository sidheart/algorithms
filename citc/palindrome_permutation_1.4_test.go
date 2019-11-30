package citc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromePermutation(t *testing.T) {
	s := "Tact Coa"
	assert.True(t, isPalindromePermutation(s))
	s = "Frosted Butts"
	assert.False(t, isPalindromePermutation(s))
	s = " "
	assert.True(t, isPalindromePermutation(s))
	s = "          a"
	assert.True(t, isPalindromePermutation(s))
}
