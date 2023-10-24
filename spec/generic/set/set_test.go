package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := New[int]()
	assert.Equal(t, 0, s.Len())

	s.Add(1)
	assert.Equal(t, 1, s.Len())

	s.Add(2)
	assert.Equal(t, 2, s.Len())

	s.Add(1)
	assert.Equal(t, 2, s.Len())

	s.Remove(1)
	assert.Equal(t, 1, s.Len())
}