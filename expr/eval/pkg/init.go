package pkg

//go generate github.com/traefik/yaegi/stdlib
//#go:generate ../internal/cmd/extract/extract sort strconv strings sync sync/atomic
//go:generate ./extract github.com/samber/lo
import "github.com/traefik/yaegi/stdlib"
var Symbols = stdlib.Symbols