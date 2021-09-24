package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	getUserEndpoint                 endpoint.Endpoint
	getUsersEndpoint                endpoint.Endpoint
	getPerpetatorEndpoint           endpoint.Endpoint
	getConnectionsEndpoint          endpoint.Endpoint
	postConnectionEndpoint          endpoint.Endpoint
	changeConnectionEndpoint        endpoint.Endpoint
	getConnectionRequestsEndpoint   endpoint.Endpoint
	decideConnectionRequestEndpoint endpoint.Endpoint
}

func NewEndpoints(s Service) Endpoints {
	return Endpoints{
		makeGetUserEndpoint(s),
		makeGetUsersEndpoint(s),
		makeGetPerpetatorEndpoint(s),
		makeGetConnectionsEndpoint(s),
		makePostConnectionsEndpoint(s),
		makeChangeConnectionsEndpoint(s),
		makeGetConnectionRequestsEndpoint(s),
		makeDecideConnectionRequestEndpoint(s),
	}
}

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getUserRequest)
		res, err := s.getUser(ctx, req)

		return res, err
	}
}

func makeGetUsersEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getUsersRequest)
		res, err := s.getUsers(ctx, req)

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

func makeGetConnectionRequestsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getConnectionRequestsRequst)
		res, err := s.getConnectionRequests(ctx, req)

		return res, err
	}
}

func makeDecideConnectionRequestEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(decideConnectionRequestRequst)
		res, err := s.decideConnectionRequest(ctx, req)

		return res, err
	}
}
