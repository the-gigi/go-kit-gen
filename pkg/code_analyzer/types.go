package code_analyzer

type Argument struct {
	Name string
	Type string
}

type ReturnValue struct {
	Name string
	Type string
}

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

type Code struct {
	Functions  []Function
	Interfaces []Interface
}
