package userservice

import (
	"cells/cell-trace/register/util"
	"context"
	"github.com/opentracing/opentracing-go"
)

var (
	ErrorUserNotFound = util.NewMyError(406, "用户不存在")
)

// Service Define a service interface
type IUserService interface {
	GetUserName(ctx context.Context, userId string) (string, error)

	UpdateUserName(ctx context.Context, userId, userName string) error

	HealthCheck() string
}

//UserService implement Service interface
type UserService struct {
}

func (s UserService) GetUserName(ctx context.Context, userId string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "serviceGetUserName")
	defer func() {
		span.LogKV("userId", userId)
		span.Finish()
	}()

	var userName string
	if value, ok := userMap[userId]; ok {
		return value, nil
	} else {
		return userName, ErrorUserNotFound
	}

}

func (s UserService) UpdateUserName(ctx context.Context, userId, userName string) error {
	if _, ok := userMap[userId]; !ok {
		return ErrorUserNotFound
	}
	userMap[userId] = userName
	return nil
}

func (s UserService) HealthCheck() string {
	return "ok"
}

//模拟数据库
var userMap map[string]string

func init() {
	userMap = make(map[string]string)
	userMap["1"] = "jack"
}

// ServiceMiddleware define service middleware
type ServiceMiddleware func(IUserService) IUserService
