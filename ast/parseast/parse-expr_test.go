package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"testing"
)

func getFuncDecl(src string, fnname string) (*token.FileSet, *ast.FuncDecl) {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic("解析源文件出错:" + err.Error())
	}
	for _, decl := range fileAst.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == fnname {
			return fset, funcDecl
		}
	}
	return fset, nil
}

type funcDeclStmt struct {
	exprs   []*ast.ExprStmt
	assigns []*ast.AssignStmt
}

func parseFuncDeclStmt(funcDecl *ast.FuncDecl) *funcDeclStmt {
	r := &funcDeclStmt{}

	for _, stmt := range funcDecl.Body.List {
		if stmt, ok := stmt.(*ast.ExprStmt); ok {
			r.exprs = append(r.exprs, stmt)
		}
		if stmt, ok := stmt.(*ast.AssignStmt); ok {
			r.assigns = append(r.assigns, stmt)
		}
	}
	return r
}

func parseStmt(stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		printExpr(stmt.X)
	case *ast.AssignStmt:
		printAssignStmt(stmt)
	}
}
func printAssignStmt(assignStmt *ast.AssignStmt) {
	println("printAssignStmt:")
	print("left:")
	for _, expr := range assignStmt.Lhs {
		printExpr(expr)
	}
	print("right:")
	for _, expr := range assignStmt.Rhs {
		printExpr(expr)
	}
	println("\n---------assignStmt end-----------------")
}

func printCallExpr(callExpr *ast.CallExpr) {
	println("printCallExpr:")
	switch fun := callExpr.Fun.(type) {
	case *ast.Ident:
		// Direct function call like: MyFunc()
		println("funcName:", fun.Name)
	case *ast.SelectorExpr:
		// Method call like: x.Method()
		objname := fun.X.(*ast.Ident).Name
		println("obj.method:", objname+"."+fun.Sel.Name+"()")
	}
	args := []string{}
	for _, arg := range callExpr.Args {
		switch arg := arg.(type) {
		case *ast.BasicLit:
			args = append(args, arg.Value)
		case *ast.Ident:
			args = append(args, arg.Name)
		}
	}
	println("args:", strings.Join(args, ","))
	println("---------funcExpr end-----------------")

}
func printExpr(expr ast.Expr) {
	switch expr := expr.(type) {
	case *ast.CallExpr:
		callExpr := expr
		printCallExpr(callExpr)
	case *ast.Ident: // r := xxx
		println("ident:", expr.Name)
	case *ast.BasicLit:
		println("basicLit:", expr.Value)
	case *ast.KeyValueExpr:
		print("\tkeyValueExpr:")
		print(expr.Key.(*ast.Ident).Name + ":")
		printExpr(expr.Value)
	case *ast.UnaryExpr: // &v, *v, +v, -v, !v, ^v, <-
		print("unaryExpr:", expr.Op.String())
		printExpr(expr.X)
	case *ast.BinaryExpr: // x + y, x - y, x * y, x / y, x % y, x & y, x | y, x ^ y, x &^ y, x << y, x >> y
		print("binaryExpr:", expr.Op.String())
		printExpr(expr.X)
		printExpr(expr.Y)
	case *ast.ParenExpr: // (expr)
		print("parenExpr:")
		printExpr(expr.X)
	case *ast.CompositeLit: // T{1, 2, 3}
		print("compositeLit{}:")
		printExpr(expr.Type)
		for _, elt := range expr.Elts {
			printExpr(elt)
		}
	}
}

// 追加ast
func TestParseCallExpr(t *testing.T) {
	const content = `
	package main
	type R struct {
        a int
        b int
    }
	func (r *R) Dump(username) {
		println(username)
	}
	func init(){
		r := &R{1,2}
		r = &R{b:2,a:3}
		r.Dump("xx")
	}
	`

	_, funcDecl := getFuncDecl(content, "init")
	stmts := parseFuncDeclStmt(funcDecl)
	for _, stmt := range stmts.exprs {
		parseStmt(stmt)
	}
	for _, stmt := range stmts.assigns {
		parseStmt(stmt)
	}
}
