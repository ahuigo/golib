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

var _ = Describe("Testing", func() {
		It("client timeout", func() {
			defer gock.Off()

			gock.New("https://baseurl.com").
				Get("/url").
				ReplyError(&timeoutError{err: "net/http: timeout awaiting response headers", timeout: true})
		})
	})

})
