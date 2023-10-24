// BEGIN: 8a7b5c3d5f3c
package lsof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLsofTcps(t *testing.T) {
	out, err := GetLsofTcps()
	assert.NoError(t, err)
	_ = out
	// assert.NotEmpty(t, out)
}
