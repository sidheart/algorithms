package dynamic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEditDistance(t *testing.T) {
	assert.Equal(t, 1, EditDistance("spice", "splice"))
	assert.Equal(t, 1, EditDistance("nice", "rice"))
	assert.Equal(t, 5, EditDistance("hypen", "hyphenated"))
	assert.Equal(t, EditDistance("bland", "bind"), 2)
}
