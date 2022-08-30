package chapter_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOneAway(t *testing.T) {
	s1, s2 := "pale", "ple"
	assert.True(t, isOneAway(s1, s2))
	s1, s2 = "pales", "pale"
	assert.True(t, isOneAway(s1, s2))
	s1, s2 = "pale", "bale"
	assert.True(t, isOneAway(s1, s2))
	s1, s2 = "pale", "bake"
	assert.False(t, isOneAway(s1, s2))
}
