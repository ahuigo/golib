package fxdemo

import (
	"fmt"
	"testing"

	"go.uber.org/fx"
)

type Connection struct{}
type Connection2 struct{}

func NewReadOnlyConnection1() (*Connection, error) {
	println("1.1 new connection")
	return &Connection{}, nil
}
func NewReadOnlyConnection2() (*Connection2, error) {
	println("1.2 new connection2")
	return &Connection2{}, nil
}

func printConn(conn *Connection) {
	fmt.Println("2. invoke conn:", conn)

}

func TestFx(t *testing.T) {
	app := fx.New(
		fx.Provide(NewReadOnlyConnection1),
		fx.Provide(NewReadOnlyConnection2),
		fx.Invoke(printConn),
	)

	fmt.Println("3. exec start...")
	_ = app
	// app.Run() //daemon
}
