package server

import (
	"fmt"
	"ginapp/test"
	"net/http"
	"testing"

	"github.com/ahuigo/requests"
)

func TestBind(t *testing.T) {
	// build request
	data := User{
		City: "Beijing",
	}
	req, _ := requests.BuildRequest("POST", "/bind?name=Alex", requests.Jsoni(data))

	// send request
	respRecorder, ctx := test.CreateTestCtx(req)
	BindServer(ctx)

	// test response status
	if respRecorder.Code != http.StatusOK {
		errors := ctx.Errors.Errors()
		fmt.Println("output", errors)
		t.Fatal("Expect code 200, but get ", respRecorder.Code, "reason", respRecorder.Body)
	}

	// test response body
	recvUser := User{}
	resp := requests.BuildResponse(respRecorder.Result())
	resp.Json(&recvUser)
	if recvUser.Name != "Alex" || recvUser.City != data.City {
		t.Fatalf("unexpected response:%v", resp.Text())
	}

}
