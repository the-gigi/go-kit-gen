package code_generator

import ()

const decodeTemplate = `
func decode{{ name }}Request(_ context.Context, r interface{}) (interface{}, error) {
	request := r.(*pb.{{ name }}Request)
	return om.{{ name }}Request{
		{{ fields }}
	}, nil
}
`

const encodeTemplate = `
func encode{{ name }}Response(_ context.Context, r interface{}) (interface{}, error) {
	return r, nil
}
`
const makeEndpoint = `
func make{{ name }}Endpoint(svc om.NewsManager) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(om.{{ name }}Request)
		r, err := svc.{{ name }}(req)
		res := New{{ name }}Response(r, err)
		if err != nil {
			res.Err = err.Error()
		}
	}
}
`

const handler = `
type handler struct {
	getNews grpctransport.Handler
}

func (s *handler) {{ name }}(ctx context.Context, r *pb.{{ name }}Request) (*pb.{{ name }}Response, error) {
	_, resp, err := s.{{ name }}.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.{{ name }}Response), nil
}
`

const handlerEndpoint = `
		{{ lowerCaseName }}: grpctransport.NewServer(
			make{{ name }}Endpoint(svc),
			decode{{ name }}Request,
			encode{{ name }}Response,
		),
`
const newServer = `
func new{{ serviceName }}Server(svc om.{{ serviceName }}Manager) pb.{{ serviceName }}Server {
	return &handler{
		{{ handlerEndpoints }}
		),
	}
}
`
