package citc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlify(t *testing.T) {
	s := "hi my name is bob"
	assert.Equal(t, "hi%20my%20name%20is%20bob", urlify(s))
	assert.Equal(t, "hi my name is bob", s)
	s = " "
	assert.Equal(t, "%20", urlify(s))
	s = "  "
	assert.Equal(t, "%20%20", urlify(s))
}
