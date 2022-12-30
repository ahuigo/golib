package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	assert.Equal(t, Length("abc"), 3)
	assert.Equal(t, LengthA([]int{1, 2, 3}), 3)
	assert.Equal(t, LengthM(map[int]int{1: 1, 2: 2}), 2)
}
