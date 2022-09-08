package main
//https://groups.google.com/g/golang-nuts/c/Nt0hVV_nqHE

import (
	"fmt"

	"github.com/jtolds/gls"
)

func main() {
	var (
		mgr            = gls.NewContextManager()
		request_id_key = gls.GenSym()
	)

	MyLog := func() {
		if request_id, ok := mgr.GetValue(request_id_key); ok {
			fmt.Println("My request id is:", request_id)
		} else {
			fmt.Println("No request id found")
		}
	}

	mgr.SetValues(gls.Values{request_id_key: "12345"}, func() {
		MyLog()
	})
	MyLog()

}
