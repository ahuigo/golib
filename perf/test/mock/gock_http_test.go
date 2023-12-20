package m

import (
	"strings"
	"testing"

	"github.com/ahuigo/requests"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestFoo(t *testing.T) {
	//defer gock.Off()

	// mock response
	gock.New("http://m.com").
		Post("/bar").
		Persist(). //Times(10)
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	// send request
	resp, err := requests.Post("http://m.com/bar")
	// 不能用: resty
	// resp, err := resty.New().R().Post("http://m.com/bar")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 200, resp.R.StatusCode)
	assert.Equal(t, `{"foo":"bar"}`, strings.TrimSpace(resp.Text()))
}
