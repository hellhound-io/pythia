package oracle

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeAdvocateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.Advocate(request.(Query))
	}
}
