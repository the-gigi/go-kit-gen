package code_analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"

	"fmt"
	"golang.org/x/tools/go/ast/inspector"

	"io/ioutil"
)

type Parser struct {
}

func (p *Parser) Parse(filename string) (result *AST, err error) {
	fset := token.NewFileSet() // positions are relative to fset
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	f, err := parser.ParseFile(fset, filename, source, 0)
	if err != nil {
		return
	}

	nodeFunc := func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
		case *ast.Ident:
			s = x.Name
		case *ast.FuncLit:
			s = x.Name

		}
		fmt.Println(s)
		return true

	}

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
	})

	return
}
