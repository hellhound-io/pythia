package node

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeFindAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return s.FindAll(ctx)
	}
}

func makeSaveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		node := request.(Node)
		return node, s.Save(&node)
	}
}
