package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strconv"
	"strings"
	"testing"
)

var srcfile = "./src.go.txt"
var outFile, _ = os.Create("tmp/a.go.txt")

const srccode = `
	package main
	func Loop(n int) { }
	func Loop2(n int) { }
	func TestX() {
		n := 10
		Loop(n)
	}
	`

// 追加ast
func TestAstAppend(t *testing.T) {

	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, srcfile, nil, parser.ParseComments)
	if err != nil {
		println("解析源文件出错:", err)
		return
	}
	ast.Print(fset, fileAst)

	// 遍历所有函数声明，找到main函数
	for _, decl := range fileAst.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && strings.HasPrefix(funcDecl.Name.Name, "Test") {
			// 在函数体的语句列表末尾添加println("main exit")语句
			stmt := &ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: ast.NewIdent("println"), // identifier
					Args: []ast.Expr{
						ast.NewIdent(`"test end"`),
						&ast.BasicLit{
							Kind:  token.INT,
							Value: "100",
						},
					},
				},
			}
			funcDecl.Body.List = append(funcDecl.Body.List, stmt)
		}
	}

	// 将修改后的AST重新打印为Go代码并输出
	err = printer.Fprint(outFile, fset, fileAst)
	if err != nil {
		println("输出代码出错:", err)
	}

}

// 替换Ast节点
func TestAstReplaceArg(t *testing.T) {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, srcfile, nil, parser.ParseComments)
	if err != nil {
		println("解析源文件出错:", err)
		return
	}

	// 遍历所有函数声明，找到main函数
	for _, decl := range fileAst.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && strings.HasPrefix(funcDecl.Name.Name, "Test") {
			// 将函数体内的调用语句Loop(n) 替换为 Loop(n+1)
			// 遍历函数体内的所有语句
			for _, stmt := range funcDecl.Body.List {
				// 查找表达式语句
				exprStmt, ok := stmt.(*ast.ExprStmt)
				if ok {
					// 查找调用表达式
					callExpr, ok := exprStmt.X.(*ast.CallExpr)
					if ok {
						// 查找函数名为Loop的调用
						funcIdent, ok := callExpr.Fun.(*ast.Ident)
						if ok && funcIdent.Name == "Loop" {
							// 假设Loop函数只有一个参数
							if len(callExpr.Args) == 1 {
								// 获取第一个参数
								switch arg := callExpr.Args[0].(type) {
								case *ast.BasicLit:
									if arg.Kind == token.INT {
										// 将参数值加1
										val, _ := strconv.Atoi(arg.Value)
										arg.Value = strconv.Itoa(val + 1)
									}
								case *ast.Ident:
									// 如果是变量，创建一个新的加法表达式
									newArg := &ast.BinaryExpr{
										X:  arg,
										Op: token.ADD,
										Y: &ast.BasicLit{
											Kind:  token.INT,
											Value: "1",
										},
									}
									callExpr.Args[0] = newArg
								}
							}
						}
					}
				}
			}
		}
	}

	// 将修改后的AST重新打印为Go代码并输出
	err = printer.Fprint(outFile, fset, fileAst)
	if err != nil {
		println("输出代码出错:", err)
	}

}

func TestAstReplaceFunc(t *testing.T) {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, "", srccode, parser.ParseComments)
	if err != nil {
		println("解析源文件出错:", err)
		return
	}

	// 遍历所有函数声明，找到main函数
	for _, decl := range fileAst.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && strings.HasPrefix(funcDecl.Name.Name, "Test") {
			// 遍历函数体内的所有语句
			for _, stmt := range funcDecl.Body.List {
				// 查找表达式语句
				exprStmt, ok := stmt.(*ast.ExprStmt)
				if !ok {
					continue
				}
				callExpr, ok := exprStmt.X.(*ast.CallExpr)
				if !ok {
					continue
				}
				// 将函数体内的调用语句Loop(n) 替换为 Loop2(n)
				funcIdent, ok := callExpr.Fun.(*ast.Ident)
				if ok && funcIdent.Name == "Loop" {
					funcIdent.Name = "Loop2"
				}
			}
		}
	}

	// 打印AST
	ast.Print(fset, fileAst)

	// 将修改后的AST重新打印为Go代码并输出
	err = printer.Fprint(outFile, fset, fileAst)
	if err != nil {
		println("输出代码出错:", err)
	}
}
