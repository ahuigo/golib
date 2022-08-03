package fxdemo

import (
	"testing"

	"go.uber.org/fx"
)

func TestOutInName(t *testing.T) {
	type result struct {
		fx.Out
		Connection1 *Connection `name:"ro"`
		Connection2 *Connection `name:"rw"`
	}
	NewReadOnlyConnection := func() (*Connection, error) {
		println("1.1 new connection")
		return &Connection{}, nil
	}

	// In
	type InData struct {
		fx.In
		RoConn *Connection `name:"ro"`
		RwConn *Connection `name:"rw"`
	}

	app := fx.New(
		fx.Provide(func() (result, error) {
			conn, _ := NewReadOnlyConnection()
			conn2, err := NewReadOnlyConnection()
			return result{Connection1: conn, Connection2: conn2}, err
		}),
		fx.Invoke(func(data InData) {
			t.Log("ro:", data.RoConn)
			t.Log("rw:", data.RwConn)
		}),
	)

	_ = app

}

// Annotated 不用构建 fx.Out objects.
func TestAnnotatedName(t *testing.T) {
	NewReadOnlyConnection := func() (*Connection, error) {
		return &Connection{}, nil
	}

	// In
	type InData struct {
		fx.In
		RoConn *Connection `name:"ro"`
		RwConn *Connection `name:"rw"`
	}

	app := fx.New(
		fx.Provide(fx.Annotated{
			Name:   "ro",
			Target: NewReadOnlyConnection,
		}),
		fx.Provide(fx.Annotated{
			Name:   "rw",
			Target: NewReadOnlyConnection,
		}),
		fx.Invoke(func(data InData) {
			t.Log("ro:", data.RoConn)
			t.Log("rw:", data.RwConn)
		}),
	)

	_ = app
}
