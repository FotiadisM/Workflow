package user

import (
	"context"
	"errors"
	"fmt"
)

type Service interface {
	getUser(ctx context.Context, req getUserRequest) (res getUserResponse, err error)
	getPerpetator(ctx context.Context, req getPerpetatorRequest) (res getPerpetatorResponse, err error)
	getConnections(ctx context.Context, req getConnectionsRequest) (res getConnectionsResponse, err error)
	postConnection(ctx context.Context, req postConnectionRequest) (res postConnectionResponse, err error)
	changeConnection(ctx context.Context, req changeConnectionRequest) (res changeConnectionResponse, err error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return service{r}
}

func (s service) getUser(ctx context.Context, req getUserRequest) (res getUserResponse, err error) {
	u, err := s.repo.GetUserByID(ctx, req.UserID)
	if err != nil {
		res.Err = err
		return
	}

	res.User = &u

	return
}

func (s service) getPerpetator(ctx context.Context, req getPerpetatorRequest) (res getPerpetatorResponse, err error) {
	u, err := s.repo.GetPerpetator(ctx, req.PerpID)
	if err != nil {
		res.Err = errors.New("failed to fetch perpetator")
		return
	}

	res.FName = u.FName
	res.LName = u.LName
	res.Company = u.Company
	res.Position = u.Position
	res.ProfilePic = u.ProfilePic

	return
}

func (s service) getConnections(ctx context.Context, req getConnectionsRequest) (res getConnectionsResponse, err error) {
	cons, err := s.repo.GetConnections(ctx, req.UserID)
	if err != nil {
		res.Err = fmt.Errorf("failed to fetch connections: %w", err)
		return
	}

	res.Connections = cons

	return
}

func (s service) postConnection(ctx context.Context, req postConnectionRequest) (res postConnectionResponse, err error) {
	id, err := s.repo.CreateConnectionRequest(ctx, req.UserID, req.User2ID)
	if err != nil {
		res.Err = errors.New("failed to create connection request")
		return
	}

	res.ConnID = id

	return
}

func (s service) changeConnection(ctx context.Context, req changeConnectionRequest) (res changeConnectionResponse, err error) {

	if req.Accept {
		err = s.repo.AcceptConnectionRequest(ctx, req.ConnID)
	} else {
		err = s.repo.RejectConnectionRequest(ctx, req.ConnID)
	}

	if err != nil {
		res.Err = err
	}

	return
}
