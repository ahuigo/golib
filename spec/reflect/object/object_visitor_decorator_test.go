package demo

import (
	"fmt"
	"testing"
)

/*
*******************************
info and visitor
*/
type VisitorFunc func(*Info, error) error
type Visitor interface {
	Visit(VisitorFunc) error
}
type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

type LogVisitor struct {
	visitor Visitor
}

func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

type NameVisitor struct {
	visitor Visitor
}

func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

func TestVisitor(t *testing.T) {
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	fmt.Printf("%#v\n", v)
	loadFile := func(info *Info, err error) error {
		fmt.Printf("execut fn...\n")
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)

}

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []WrapVisitorFunc
}

type WrapVisitorFunc func(VisitorFunc) VisitorFunc

func wrapNameVisitor(fn VisitorFunc) VisitorFunc {
	return func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		fmt.Println("NameVisitor() after call function")
		return err
	}
}
func wrapLogVisitor(fn VisitorFunc) VisitorFunc {
	return func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	for _, wrapFn := range v.decorators {
		fn = wrapFn(fn)
	}
	return v.visitor.Visit(fn)
}

func TestDecorators(t *testing.T) {
	NewDecoratedVisitor := func(v Visitor, fn ...WrapVisitorFunc) Visitor {
		if len(fn) == 0 {
			return v
		}
		return DecoratedVisitor{v, fn}
	}

	info := Info{}
	var v Visitor = &info
	v = NewDecoratedVisitor(v, wrapLogVisitor, wrapNameVisitor)

	loadFile := func(info *Info, err error) error {
		fmt.Printf("execut fn...\n")
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	v.Visit(loadFile)
}
