package userservice

import (
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
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
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
		UserId: userId,
		UserName:userName,
	}, nil
}

func encodeUpdateUserNameResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
