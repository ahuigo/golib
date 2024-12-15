package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAstWalk(t *testing.T) {
	const srccode = `
	package mainx
	import "fmt"
	func Loop4(n int) { }
	`
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, "", srccode, parser.ParseComments)
	if err != nil {
		println("解析源文件出错:", err)
		return
	}

	// Walk the AST.
	ast.Inspect(fileAst, func(n ast.Node) bool {
		// Called recursively.
		ast.Print(fset, n)
		return true
	})
}
