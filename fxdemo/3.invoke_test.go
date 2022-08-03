package fxdemo

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/fx"
)

type Boy struct {
	Name string
}

type Girl struct {
	Name string
}

type Class struct {
	b *Boy
	g *Girl
}

// InitBoy 构造函数
func InitBoy() *Boy {
	fmt.Println("1.1 init: Boy")
	return &Boy{Name: "xiaoming"}
}

// InitGirl 构造函数
func InitGirl() *Girl {
	fmt.Println("1.2 init: Girl")
	return &Girl{Name: "xiaohong"}
}

// InitClass 构造函数
func InitClass(b *Boy, g *Girl) *Class {
	fmt.Println("1.3 Init: Class<-(Boy, Girl)")
	c := &Class{
		b: b,
		g: g,
	}
	return c
}

func TestInvoke(t *testing.T) {
	// 创建一个container
	app := fx.New(
		// 注入构造函数
		fx.Provide(InitGirl, InitBoy),
		// 执行：自动调用构造
		fx.Invoke(InitClass),
	)

	fmt.Println("2. begin Start")
	err := app.Start(context.Background())
	fmt.Println("3. start end")

	if err != nil {
		fmt.Println("err:", err)
	}

}
