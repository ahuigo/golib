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
		Cors *Cors `name:"cors"` // 声明要调用的数据标签：name:"cors"
	}
	CORS := func() *Cors {
		return &Cors{}
	}
	module1 := fx.Module(
		"mod1",
		fx.Provide(
			fx.Annotate(CORS, fx.ResultTags(`name:"cors"`)), // 声明要注入数据标签: name:"cors"
		),
	)
	module2 := fx.Module("api", // module　将多个module+provide 合并成module
		module1,
	)
	moduleAll := fx.Options( // 也可以用Options, 不用ModuleName
		module2,
		//...
	)

	// test
	// t.Run("Provide", func(t *testing.T) {
	// 	t.Parallel()
	app := fx.New(
		moduleAll,
		fx.Invoke(func(i In) {
			assert.Equal(t, false, i.Cors == nil)
		}),
	)
	// app.Start(context.Background())
	_ = app
	// })

}
