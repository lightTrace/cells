package userservice

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

var (
	ErrInvalidRequestType = errors.New("RequestType has only four type: Add,Subtract,Multiply,Divide")
)

type Endpoints struct {
	GetUserNameEndpoint    endpoint.Endpoint
	UpdateUserNameEndpoint endpoint.Endpoint
}

type GetUserNameRequest struct {
	UserId string `json:"userId"`
}

type GetUserNameResponse struct {
	UserName string `json:"userName"`
}

type UpdateUserNameRequest struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

type UpdateUserNameResponse struct {
}

func MakeGetUserNameEndpoint(svc IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserNameRequest)
		userName, err := svc.GetUserName(req.UserId)
		if err != nil {
			return nil, err
		}
		return GetUserNameResponse{
			UserName: userName,
		}, nil

	}
}

func MakeUpdateUserNameEndpoint(svc IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateUserNameRequest)
		return nil, svc.UpdateUserName(req.UserId, req.UserName)
	}
}
