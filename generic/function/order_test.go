package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert.True(t, Equal(1)(1))
	assert.False(t, Equal(1)(2))
	assert.True(t, Equal("1")("1"))
	assert.False(t, Equal("1")("2"))
}
