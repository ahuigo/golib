package main

import (
	"testing"

	"github.com/ahuigo/requests"
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
	if err != nil {
		t.Fatal(err)
	}
    t.Log(resp.R.StatusCode)
	t.Log(resp.Text())
}
