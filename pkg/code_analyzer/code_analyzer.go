package code_analyzer

import (
	"go/ast"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/inspector"
	"fmt"

	"io/ioutil"
)


type Method {
	Name string
	Arguments
}

type Code struct {
	InterfaceName string


}

func


