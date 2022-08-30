package chapter_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnimalShelter(t *testing.T) {
	a := NewAnimalShelter()
	a.Enqueue(Cat{1})
	a.Enqueue(Dog{2000})
	assert.Equal(t, Cat{1}, a.DequeueAny())
	a.Enqueue(Cat{3000})
	assert.Equal(t, Dog{2000}, a.DequeueAny())
	assert.Equal(t, Cat{3000}, a.DequeueAny())
	assert.True(t, a.IsEmpty())
	a.Enqueue(Cat{5})
	a.Enqueue(Dog{6})
	a.Enqueue(Dog{7})
	assert.Equal(t, Cat{5}, a.DequeueCat())
	assert.Equal(t, Dog{6}, a.DequeueDog())
}
