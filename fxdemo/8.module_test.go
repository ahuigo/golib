package fxdemo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

type Cors struct{}

func TestModule(t *testing.T) {
	type In struct {
		fx.In
		Cors *Cors `name:"cors"`
	}
	CORS := func() *Cors {
		return &Cors{}
	}
	module1 := fx.Module(
		"mod1",
		fx.Provide(
			fx.Annotate(CORS, fx.ResultTags(`name:"cors"`)),
		),
	)
	module2 := fx.Module("api",
		module1,
	)
	_ = module2

	// test
	// t.Run("Provide", func(t *testing.T) {
	// 	t.Parallel()
	app := fx.New(
		module1,
		fx.Invoke(func(i In) {
			assert.Equal(t, false, i.Cors == nil)
		}),
	)
	// app.Start(context.Background())
	_ = app
	// })

}
