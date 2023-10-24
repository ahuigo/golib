// BEGIN: 8f7e3d5b0f5c
package netstat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTcpConnections(t *testing.T) {
	tcps, err := GetAllTcpConnections()
	assert.NoError(t, err)
	for _, tcp := range tcps {
		t.Logf("%+v", tcp)
	}

}
