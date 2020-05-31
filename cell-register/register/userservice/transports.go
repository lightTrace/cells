package userservice

import (
	"cells/cell-register/register/util"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoints Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		//自定义的错误解码器
		kithttp.ServerErrorEncoder(util.MyErrorEncoder),
	}

	r.Methods("GET").Path(`/user/{userId}`).Handler(kithttp.NewServer(
		endpoints.GetUserNameEndpoint,
		decodeGetUserNameRequest,
		encodeGetUserNameResponse,
		options...,
	))
	r.Methods("POST").Path(`/user/{userId}/{userName}`).Handler(kithttp.NewServer(
		endpoints.UpdateUserNameEndpoint,
		decodeUpdateUserNameRequest,
		encodeUpdateUserNameResponse,
		options...,
	))
	// create health check handler
	r.Methods("GET").Path("/health").Handler(kithttp.NewServer(
		endpoints.HealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeArithmeticResponse,
		options...,
	))

	return r
}

func decodeGetUserNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return nil, ErrorBadRequest
	}

	return GetUserNameRequest{
		UserId: userId,
	}, nil
}

func encodeGetUserNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeUpdateUserNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return nil, ErrorBadRequest
	}
	userName, ok := vars["userName"]
	if !ok {
		return nil, ErrorBadRequest
	}

	return UpdateUserNameRequest{
		UserId:   userId,
		UserName: userName,
	}, nil
}

func encodeUpdateUserNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encodeArithmeticResponse encode response to return
func encodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// decodeHealthCheckRequest decode request
func decodeHealthCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return HealthRequest{}, nil
}
