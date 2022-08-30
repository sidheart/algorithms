package dynamic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibMemo(t *testing.T) {
	assert.Equal(t, Fib(0), 0)
	assert.Equal(t, Fib(1), 1)
	assert.Equal(t, Fib(2), 1)
	assert.Equal(t, Fib(68), 72723460248141)
}
