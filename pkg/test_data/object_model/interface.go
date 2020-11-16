package object_model

type Op1Request struct {
	A int
	B int
}

type Op1Response struct {
	Result int
}

type Op2Request struct {
	X map[string]interface{}
	Y *[]bool
}

type Op2Response struct {
	Z int
}

type Op3Request struct {
}

type Op3Response struct {
}

type Foo interface {
	Op1(r *Op1Request) (*Op1Response, error)
	Op2(r *Op2Request) (*Op2Response, error)
	Op3(r *Op3Request) (*Op3Response, error)
}
