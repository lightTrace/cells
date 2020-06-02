package userservice

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"log"
)

var (
	ErrInvalidRequestType = errors.New("RequestType has only four type: Add,Subtract,Multiply,Divide")
)

type Endpoints struct {
	GetUserNameEndpoint    endpoint.Endpoint
	UpdateUserNameEndpoint endpoint.Endpoint
	HealthCheckEndpoint    endpoint.Endpoint
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
		userName, err := svc.GetUserName(ctx, req.UserId)
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
		return nil, svc.UpdateUserName(ctx, req.UserId, req.UserName)
	}
}

// HealthRequest 健康检查请求结构
type HealthRequest struct{}

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status string `json:"status"`
}

// MakeHealthCheckEndpoint 创建健康检查Endpoint
func MakeHealthCheckEndpoint(svc IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		log.Print("health check")
		return HealthResponse{status}, nil
	}
}
