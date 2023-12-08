package router

import (
	"fmt"
	"net/http"
	"testing"

	"ginapp/test"

	"github.com/ahuigo/requests"
)

func TestAddUser(t *testing.T) {
	// build request
	data := User{
		Name: "testuser",
		City: "Beijing",
	}
	req, _ := requests.BuildRequest("POST", "/api/v1/user", requests.Jsoni(data))

	// send request
	respRecorder, ctx := test.CreateTestCtx(req)
	AddUser(ctx)

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
	if recvUser.Name != data.Name || recvUser.City != data.City {
		t.Fatalf("unexpected response:%v", resp.Text())
	}

}
