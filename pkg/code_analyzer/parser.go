package code_analyzer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"

	"github.com/pkg/errors"
)

type Parser struct {
}

func printFunc(name string, f *ast.FuncType) {
	params := f.Params.List
	parts := []string{"func ", name, "("}
	fmt.Println("name:", name)
	fmt.Println("params:")
	for i, p := range params {
		paramType := fmt.Sprintf("%v", p.Type)
		if i > 0 {
			parts = append(parts, ", ")
		}
		parts = append(parts, p.Names[0].String(), " ", paramType)
	}
	parts = append(parts, ")")
	appendResult := func(r *ast.Field) {
		resultType := fmt.Sprintf("%v", r.Type)
		if len(r.Names) > 0 {
			parts = append(parts, " (", r.Names[0].Name, " ", resultType, ")")
		} else {
			parts = append(parts, resultType)
		}
	}

	if f.Results != nil {
		results := f.Results.List
		if len(results) == 1 {
			parts = append(parts, " ")
			appendResult(results[0])
		} else if len(results) > 1 {
			parts = append(parts, " (")
			for _, r := range results {
				appendResult(r)
			}
			parts = append(parts, ")")
		}
	}

	s := strings.Join(parts, "")
	fmt.Println(s)
}

func parseCallable(name string, f *ast.FuncType) string {
	params := f.Params.List
	parts := []string{"func ", name, "("}
	fmt.Println("name:", name)
	fmt.Println("params:")
	for i, p := range params {
		paramType := fmt.Sprintf("%v", p.Type)
		if i > 0 {
			parts = append(parts, ", ")
		}
		parts = append(parts, p.Names[0].String(), " ", paramType)
	}
	parts = append(parts, ")")
	appendResult := func(r *ast.Field) {
		resultType := fmt.Sprintf("%v", r.Type)
		if len(r.Names) > 0 {
			parts = append(parts, " (", r.Names[0].Name, " ", resultType, ")")
		} else {
			parts = append(parts, resultType)
		}
	}

	results := f.Results.List
	if len(results) == 1 {
		parts = append(parts, " ")
		appendResult(results[0])
	} else if len(results) > 1 {
		parts = append(parts, " (")
		for i, r := range results {
			if i > 0 {
				parts = append(parts, ", ")
			}
			appendResult(r)
		}
		parts = append(parts, ")")
	}

	return strings.Join(parts, "")
}

func parseInterface(name string, it *ast.InterfaceType) (result string) {
	fmt.Println(name, "%v", it)

	lines := []string{"type " + name + " interface", "{"}

	methods := it.Methods.List
	for _, m := range methods {
		name := m.Names[0].String()
		funcType := m.Type.(*ast.FuncType)
		line := parseCallable(name, funcType)
		lines = append(lines, "\t"+line)
	}

	lines = append(lines, "}", "")

	result = strings.Join(lines, "\n")
	return
}

func nodeFunc(n ast.Node) bool {
	if n == nil {
		return false
	}
	switch x := n.(type) {
	case *ast.ArrayType:
		//fmt.Println("%v", x)
	case *ast.FuncDecl:
		printFunc(x.Name.String(), x.Type)

	case *ast.FuncType:
		printFunc("func", x)
		//result := parseCallable()
		//fmt.Println(result)

	case *ast.TypeSpec:
		specType := x.Type
		switch t := specType.(type) {
		case *ast.InterfaceType:
			name := x.Name.String()
			result := parseInterface(name, t)
			fmt.Println(result)
		}
	default:
		fmt.Printf("%v\n", x)
		fmt.Println("--------------")
	}

	return true
}

func (p *Parser) Parse(filename string) (result *Code, err error) {
	fset := token.NewFileSet() // positions are relative to fset
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	f, err := parser.ParseFile(fset, filename, source, 0)
	if err != nil {
		return
	}

	// Inspect the AST and print all identifiers and literals.
	result = &Code{}
	//ast.Inspect(f, nodeFunc)
	err = p.parseFile(f, result)
	return
}

func (p *Parser) parseFile(f *ast.File, code *Code) (err error) {
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		switch x := n.(type) {
		case *ast.ArrayType:
			//fmt.Println("%v", x)
		case *ast.FuncDecl:
			//printFunc(x.Name.String(), x.Type)
			name := x.Name.String()
			t := x.Type
			result := p.parseFunction(name, nil, t.Params, t.Results)
			code.Functions = append(code.Functions, result)

		//case *ast.FuncType:
		//	printFunc("func", x)

		case *ast.TypeSpec:
			specType := x.Type
			switch t := specType.(type) {
			case *ast.InterfaceType:
				name := x.Name.String()
				var result Interface
				result, err = p.parseInterface(name, t)
				code.Interfaces = append(code.Interfaces, result)
				//fmt.Println(result)
			}
		default:
			//fmt.Printf("%v\n", x)
			//fmt.Println("--------------")
		}

		return true
	})

	return
}

func (p *Parser) parseFunction(name string,
	receiver *ast.Field,
	params *ast.FieldList,
	results *ast.FieldList) (result Function) {
	result.Name = name
	for _, p := range params.List {
		argType := fmt.Sprintf("%v", p.Type)
		arg := Argument{
			Name: p.Names[0].String(),
			Type: argType,
		}
		result.Arguments = append(result.Arguments, arg)
	}

	if results == nil {
		return
	}

	for _, r := range results.List {
		returnValue := ReturnValue{
			Type: fmt.Sprintf("%v", r.Type),
		}
		if len(r.Names) > 0 {
			returnValue.Name = r.Names[0].Name
		}
		result.Result = append(result.Result, returnValue)
	}
	return
}

func (p *Parser) parseCallable(callable *ast.Field) (result Function) {
	name := callable.Names[0].Name
	ft := callable.Type.(*ast.FuncType)
	result = p.parseFunction(name, nil, ft.Params, ft.Results)
	return
}

func (p *Parser) parseInterface(name string, it *ast.InterfaceType) (result Interface, err error) {
	if it == nil {
		err = errors.New("missing interface type")
		return
	}
	result.Name = name
	methods := it.Methods.List
	for _, m := range methods {
		f := p.parseCallable(m)
		result.Methods = append(result.Methods, f)
	}

	return
}
