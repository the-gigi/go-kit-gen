package code_analyzer

type Param struct {
	Name string
	Type string
}

type Method struct {
	Name      string
	Arguments []interface{}
	Result    []interface{}
}

type Code struct {
	InterfaceName string
	Methods       []Method
}
