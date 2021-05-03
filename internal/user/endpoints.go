package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getUserEndpoint          endpoint.Endpoint
	getPerpetatorEndpoint    endpoint.Endpoint
	getConnectionsEndpoint   endpoint.Endpoint
	postConnectionEndpoint   endpoint.Endpoint
	changeConnectionEndpoint endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetUserEndpoint(s),
		makeGetPerpetatorEndpoint(s),
		makeGetConnectionsEndpoint(s),
		makePostConnectionsEndpoint(s),
		makeChangeConnectionsEndpoint(s),
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getUserRequest)
		res, err := s.getUser(ctx, req)

		return res, err
	}
}

func makeGetPerpetatorEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getPerpetatorRequest)
		res, err := s.getPerpetator(ctx, req)

		return res, err
	}
}

func makeGetConnectionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getConnectionsRequest)
		res, err := s.getConnections(ctx, req)

		return res, err
	}
}

func makePostConnectionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postConnectionRequest)
		res, err := s.postConnection(ctx, req)

		return res, err
	}
}

func makeChangeConnectionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(changeConnectionRequest)
		res, err := s.changeConnection(ctx, req)

		return res, err
	}
}
