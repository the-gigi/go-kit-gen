// SimpleParser parses a single interface from a file
//
package code_analyzer

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

const (
	methodGroupCount      = 3
	interfaceStartPattern = `type (.*) interface \{`
	interfaceEndPattern   = `\}`
	methodPattern         = `\s(.*)\(r \*(.*)Request\) \(\*(.*)Response, error\)`
)

type SimpleParser struct {
	interfaceStartRegex *regexp.Regexp
	interfaceEndRegex   *regexp.Regexp
	methodRegex         *regexp.Regexp
}

func (p *SimpleParser) Parse(filename string) (result *Interface, err error) {
	bytes := []byte{}
	bytes, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	result, err = p.parse(string(bytes))
	return
}

//// parseMethod parses a code line that represents an interface method.
////
//// The expected format is `<MethodName>(r *<MethodName>Request) (*<MethodName>Response, error)`
////
//// For example: Op1(r *Op1Request) (*Op1Response, error)
//func (p *SimpleParser) parseMethod(line string) (result *Function, err error) {
//
//}

type ParseResult struct {
	Interface Interface
	Structs   []Struct
}

func (p *SimpleParser) parse(s string) (result *Interface, err error) {
	result = &Interface{}
	lines := strings.Split(s, "\n")
	// search for interface
	startIndex := -1
	for i, line := range lines {
		matches := p.interfaceStartRegex.FindStringSubmatch(line)
		if len(matches) > 0 {
			startIndex = i
			result = &Interface{
				Name:    matches[1],
				Methods: []Function{},
			}
			break
		}
	}

	// search for closing bracket of interface
	for _, line := range lines[startIndex+1:] {
		if p.interfaceEndRegex.MatchString(line) {
			break
		}

		// search for methods
		matches := p.methodRegex.FindStringSubmatch(line)

		// Ignore incompatible lines (could be empty or comment lines)
		if len(matches) != methodGroupCount+1 {
			continue
		}

		// Verify format
		if matches[1] != matches[2] || matches[1] != matches[3] {
			err = errors.New("Invalid method format:" + line)
		}

		name := matches[1]
		method := Function{
			Name: matches[1],
			Arguments: []Argument{
				{
					Name: "r",
					Type: "*" + name + "Request",
				},
			},
			Result: []ReturnValue{
				{
					Name: "",
					Type: "*" + name + "Response",
				},
				{
					Name: "",
					Type: "error",
				},
			},
		}
		result.Methods = append(result.Methods, method)
	}

	return
}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{
		interfaceStartRegex: regexp.MustCompile(interfaceStartPattern),
		interfaceEndRegex:   regexp.MustCompile(interfaceEndPattern),
		methodRegex:         regexp.MustCompile(methodPattern),
	}
}
