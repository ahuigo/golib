generate:
	go generate ./pkg
	# go:generate ./pkg/extract github.com/samber/lo
install:
	//go:generate ../internal/cmd/extract/extract sort strconv strings sync sync/atomic
	go build -o pkg/extract ~/go/pkg/mod/github.com/traefik/yaegi@v0.15.1/internal/cmd/extract/
