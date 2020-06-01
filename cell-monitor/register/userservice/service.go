package userservice

import "cells/cell-monitor/register/util"

var (
	ErrorUserNotFound = util.NewMyError(406, "用户不存在")
)

// Service Define a service interface
type IUserService interface {
	GetUserName(userId string) (string, error)

	UpdateUserName(userId, userName string) error

	HealthCheck() string
}

//UserService implement Service interface
type UserService struct {
}

func (s UserService) GetUserName(userId string) (string, error) {
	var userName string
	if value, ok := userMap[userId]; ok {
		return value, nil
	} else {
		return userName, ErrorUserNotFound
	}
}

func (s UserService) UpdateUserName(userId, userName string) error {
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
