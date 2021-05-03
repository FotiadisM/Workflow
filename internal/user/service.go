package user

import "context"

type Service interface {
	getUser(ctx context.Context, req getUserRequest) (res getUserResponse, err error)
	getPerpetator(ctx context.Context, req getPerpetatorRequest) (res getPerpetatorResponse, err error)
	getConnections(ctx context.Context, req getConnectionsRequest) (res getConnectionsResponse, err error)
	postConnection(ctx context.Context, req postConnectionRequest) (res postConnectionResponse, err error)
	changeConnection(ctx context.Context, req changeConnectionRequest) (res changeConnectionResponse, err error)
}

type service struct {
}

func NewService() Service {
	return service{}
}

func (s service) getUser(ctx context.Context, req getUserRequest) (res getUserResponse, err error) {
	return
}

func (s service) getPerpetator(ctx context.Context, req getPerpetatorRequest) (res getPerpetatorResponse, err error) {
	return
}

func (s service) getConnections(ctx context.Context, req getConnectionsRequest) (res getConnectionsResponse, err error) {
	return
}

func (s service) postConnection(ctx context.Context, req postConnectionRequest) (res postConnectionResponse, err error) {
	return
}

func (s service) changeConnection(ctx context.Context, req changeConnectionRequest) (res changeConnectionResponse, err error) {
	return
}
