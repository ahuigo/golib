package fxdemo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func NewForTest(tb testing.TB, opts ...fx.Option) *fx.App {
	testOpts := []fx.Option{
		// Provide both: Logger and WithLogger so that if the test
		// WithLogger fails, we don't pollute stderr.
		fx.Logger(fxtest.NewTestPrinter(tb)),
		fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(tb) }),
	}
	opts = append(testOpts, opts...)

	return fx.New(opts...)
}

/**
Annotate lets you annotate a function's parameters and returns without you having to declare separate struct definitions for them.

    func Annotate(t interface{}, anns ...Annotation) interface{}
*/
type Conn struct {
}

type asStringer struct {
	name string
}

func (as *asStringer) String() string {
	return as.name
}

type in struct {
	fx.In
	S fmt.Stringer `name:"goodStringer"`
	// S *asStringer //  `name:"goodStringer"`
}

func TestAnnotateResultName(t *testing.T) {
	tt := struct {
		provide fx.Option
		invoke  interface{}
	}{
		provide: fx.Provide(
			fx.Annotate(
				func() *asStringer {
					return &asStringer{name: "stringer"}
				},
				fx.ResultTags(`name:"goodStringer"`),
				fx.As(new(fmt.Stringer)),
			), //type: fmt.Stringer[name="goodStringer"]
		),
		invoke: func(i in) {
			assert.Equal(t, "stringer", i.S.String())
		},
	}
	app := fx.New(
		tt.provide,
		fx.Invoke(tt.invoke),
	)

	_ = app

}
func TestAnnotateParamOptional(t *testing.T) {
	type a struct{}
	type b struct{ a *a }
	type c struct{ b *b }
	newA := func() *a {
		return &a{}
	}
	newB := func(a *a, a2 *a) *b {
		return &b{a}
	}
	newC := func(b *b) *c {
		return &c{b}
	}
	app := fxtest.New(t,
		fx.Provide(
			fx.Annotate(
				newA,
				fx.ResultTags(`name:"arge2"`),
			),
			fx.Annotate(newB, fx.ParamTags(`name:"arg1" optional:"true"`, `name:"arge2"`)), // b{a:nil}
		),
		fx.Invoke(newC),
	)
	defer app.RequireStart().RequireStop()
	assert.Equal(t, app.Err(), nil)
}

/**
example:

	type Gateway struct {
		ro *Conn
		rw *Conn
	}
	NewGateway := func(ro, rw *Conn) *Gateway {
		return &Gateway{ro, rw}
	}
	app := fx.New(
		fx.Provide(
			fx.Annotate(
				NewGateway,
				fx.ParamTags(`name:"ro" optional:"true"`, `name:"rw"`),
				fx.ResultTags(`name:"foo"`),
			),
		),
	)
	defer app.RequireStart().RequireStop()

Is some like equivalent to,

    type params struct {
      fx.In

      RO *db.Conn `name:"ro" optional:"true"`
      RW *db.Conn `name:"rw"`
    }

    type result struct {
      fx.Out
      GW *Gateway `name:"foo"`
    }

     fx.Provide(func(p params) result {
       return result{GW: NewGateway(p.RO, p.RW)}
     })
**/
func TestAnnotateParamTag(t *testing.T) {
	type b struct{}
	type a struct{ b *b }
	type sliceA struct{ sa []*a }
	newA := func(b *b, b2 *b) *a {
		return &a{b}
	}
	newSliceA := func(sa []*a) *sliceA {
		return &sliceA{sa}
	}
	var got *sliceA
	fxtest.New(t,
		fx.Provide(
			fx.Annotate(
				newA,
				fx.ParamTags(`name:"arg1" optional:"true"`, `name:"arge2" optional:"true"`),
				fx.ResultTags(`group:"as"`),
			),
			fx.Annotate(
				newA,
				fx.ParamTags(`name:"arg1" optional:"true"`, `name:"arge2" optional:"true"`),
				fx.ResultTags(`group:"as"`),
			),
			// fx.Annotated{Group: "as", Target: newA},
			// fx.Annotated{Group: "as", Target: newA},
			fx.Annotate(newSliceA, fx.ParamTags(`group:"as"`)),
		),
		fx.Populate(&got),
	)
	assert.Equal(t, len(got.sa), 2)

}

func TestAnnotateAs(t *testing.T) {
	newAsStringer := func() *asStringer {
		return &asStringer{
			name: "a good stringer",
		}
	}

	type myStringer interface {
		String() string
	}
	type I interface{}
	type J interface{}

	app := fx.New(
		fx.Provide(
			fx.Annotate(newAsStringer, fx.As(new(fmt.Stringer))),
			fx.Annotate(newAsStringer, fx.As(new(myStringer))),
			fx.Annotate(
				func() (string, error) {
					return "str", nil
				},
				fx.As(new(I)), fx.As(new(J)),
			),
		),
		fx.Invoke(func(s fmt.Stringer, s2 myStringer) {
			assert.Equal(t, s.String(), "a good stringer")
			assert.Equal(t, s2.String(), "a good stringer")
		}),
		fx.Invoke(func(a I, b J) {
			assert.Equal(t, "str", a)
			assert.Equal(t, "str", b)
		}),
	)
	_ = app

}
