package code_analyzer

type Field struct {
	Name string
	Type string
}

type Argument = Field
type ReturnValue = Field

type Function struct {
	Name      string
	Arguments []Argument
	Result    []ReturnValue
}

type Method struct {
	Function
	Receiver string
}

type Interface struct {
	Name    string
	Methods []Function
}

type Struct struct {
	Name       string
	FieldNames []string
}
type Code struct {
	Functions  []Function
	Interfaces []Interface
}
