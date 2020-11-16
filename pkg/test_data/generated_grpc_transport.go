package test_data

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"

	om "github.com/the-gigi/go-kit-gen/pkg/test_data/object_model"
)

type op1Request struct {
	A int `json:"a"`
	B int `json:"b"`

	Err string
}

type Op1Response struct {
	Result int `json:"result"`
	Err string `json:"err"`
}

func decodeOp1Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request om.Op1Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}

func decodeOp2Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request om.Op2Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func decodeOp3Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request om.Op3Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func makeOp1Endpoint(svc om.Foo) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*om.Op1Request)
		if !ok {
			return nil, errors.New("invalid request")
		}
		r, err := svc.Op1(req)

		var res Op1Response
		if err != nil {
			res.Err = err.Error()
		}

		if r != nil {
			res.Result = r.Result
		}

		return res, nil
	}
}
