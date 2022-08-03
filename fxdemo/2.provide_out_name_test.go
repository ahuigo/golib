package fxdemo

import (
	"testing"

	"go.uber.org/dig"
)

// refer to : https://www.cnblogs.com/li-peng/p/14708132.html
/*
如果Provide里提供的函数，有多个函数返回的数据类型是一样的怎么处理？比如，我们的数据库有主从两个连接库，怎么进行区分？
dig可以将Provide命名以进行区分
*/
func TestProvideNameWrap(t *testing.T) {
	type DSN struct {
		Addr string
	}
	c := dig.New()

	// provide name 1
	p1 := func() (*DSN, error) {
		return &DSN{Addr: "primary DSN"}, nil
	}
	if err := c.Provide(p1, dig.Name("primary")); err != nil {
		t.Fatal(err)
	}

	// provide name 2
	p2 := func() (*DSN, error) {
		return &DSN{Addr: "secondary DSN"}, nil
	}
	if err := c.Provide(p2, dig.Name("secondary")); err != nil {
		t.Fatal(err)
	}

	// invoke name: dig.In
	type DBInfo struct {
		dig.In
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}

	if err := c.Invoke(func(db DBInfo) {
		t.Log(db.PrimaryDSN)
		t.Log(db.SecondaryDSN)
	}); err != nil {
		t.Fatal(err)
	}
}

/**
更通用的方法：dig.Out + Name
一般我们是有一个结构体来实现，dig也有相应的支持，用一个结构体嵌入dig.out来实现，
相同类型的字段在tag里设置不同的name来实现
*/
func TestProvideNameTags(t *testing.T) {
	type DSN struct {
		Addr string
	}
	c := dig.New()

	// provide multi instance with dig.Out
	type DSNRev struct {
		dig.Out
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}
	p := func() (DSNRev, error) {
		return DSNRev{
			PrimaryDSN:   &DSN{Addr: "Primary DSN"},
			SecondaryDSN: &DSN{Addr: "Secondary DSN"},
		}, nil
	}

	if err := c.Provide(p); err != nil {
		t.Fatal(err)
	}

	// invoke multiple instance with dig.In
	type DBInfo struct {
		dig.In
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}
	inv1 := func(db DBInfo) {
		t.Log(db.PrimaryDSN)
		t.Log(db.SecondaryDSN)
	}

	if err := c.Invoke(inv1); err != nil {
		t.Fatal(err)
	}
}
