package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

// 追加ast
func TestParseAllVars(t *testing.T) {
	var srcfile = "./src.go.txt"
	// var outFile, _ = os.Create("tmp/a.go.txt")

	const content = `
	package main
	var A = 10
	var B = 10
	func Loop(n int) { }
	func Loop2(n int) { }
	func TestX() {
		n := 10
		Loop(n)
	}
	`

	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, srcfile, content, parser.ParseComments)
	if err != nil {
		t.Fatal("解析源文件出错:", err)
		return
	}
	// ast.Print(fset, fileAst)

	var results []string
	// 遍历所有声明
	for _, decl := range fileAst.Decls {
		// 类型断言为 GenDecl（通用声明）
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			// 检查是否为变量声明
			if genDecl.Tok == token.VAR {
				// 获取声明的起始位置和结束位置
				start := fset.Position(genDecl.Pos())
				end := fset.Position(genDecl.End())

				// 提取变量声明的源代码
				lines := strings.Split(string(content), "\n")
				for i := start.Line - 1; i < end.Line; i++ {
					results = append(results, lines[i])
				}
			}
		}
	}

	// 将修改后的AST重新打印为Go代码并输出
	t.Log(strings.Join(results, "\n"))

}
