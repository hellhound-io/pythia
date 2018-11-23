package cryptosystem

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeFindEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.Find(request.(string))
	}
}

func makeStoreEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.Store(request.(CryptoSystem))
	}
}

func makeFindAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return s.FindAll(ctx)
	}
}
