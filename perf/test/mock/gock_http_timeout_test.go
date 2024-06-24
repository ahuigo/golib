package m

import (
	"testing"

	"github.com/ahuigo/requests"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

type timeoutError struct {
	err     string
	timeout bool
}

func (e *timeoutError) Error() string {
	return e.err
}
func (e *timeoutError) Timeout() bool {
	return e.timeout
}

func (e *timeoutError) Temporary() bool {
	return true
}

func TestTimeoutResty(t *testing.T) {
	/**
	  resty 没有用 http.DefaultTransport 被换了
	  hc.Transport = createTransport(nil)
	*/
	defer gock.Off()
	gock.New("http://m.com").
		Get("/url").
		ReplyError(&timeoutError{err: "net/http: timeout awaiting response headers", timeout: true})

	resp, err := resty.New().R().Get("http://m.com/url")
	assert.ErrorContains(t, err, "net/http: timeout awaiting response headers")
	t.Log(err, resp.IsError())

}

func TestTimeoutRequests(t *testing.T) {
	defer gock.Off()
	gock.New("http://m.com").
		Get("/url").
		ReplyError(&timeoutError{err: "net/http: timeout awaiting response headers", timeout: true})

	_, err := requests.Get("http://m.com/url")
	assert.ErrorContains(t, err, "net/http: timeout awaiting response headers")

}

func TestTimeoutRequestsRetry(t *testing.T) {
	defer gock.Off()
	// 只能触发一次
	gock.New("http://m.com").
		Get("/url").Persist().
		ReplyError(&timeoutError{err: "net/http: timeout awaiting response headers", timeout: true})

	r := requests.R().
		SetRetryCount(3).
		SetRetryCondition(func(resp *requests.Response, err error) bool {
			// return false
			return err != nil
		})
	_, err := r.Get("http://m.com/url")
	assert.ErrorContains(t, err, "net/http: timeout awaiting response headers")

}
