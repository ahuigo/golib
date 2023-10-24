package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	assert.True(t, And(true)(true))
	assert.False(t, And(true)(false))
	assert.False(t, And(false)(true))
	assert.False(t, And(false)(false))
}

func TestOr(t *testing.T) {
	assert.True(t, Or(true)(true))
	assert.True(t, Or(true)(false))
	assert.True(t, Or(false)(true))
	assert.False(t, Or(false)(false))
}

func TestNot(t *testing.T) {
	assert.False(t, Not(true))
	assert.True(t, Not(false))
}
