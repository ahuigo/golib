module m

go 1.21.4

require (
	github.com/ahuigo/requests v1.0.19
	github.com/go-resty/resty/v2 v2.10.0
	github.com/ovechkin-dm/mockio v0.4.6
	gopkg.in/h2non/gock.v1 v1.1.2
)
replace (
        github.com/go-resty/resty/v2 v2.10.0 => ./resty
)
require (
	github.com/alessio/shellescape v1.4.1 // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/ovechkin-dm/go-dyno v0.0.21 // indirect
	github.com/petermattis/goid v0.0.0-20230904192822-1876fd5063bc // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.17.0 // indirect
)
