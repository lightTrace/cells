package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func MakeHttpHandler(endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path(`/user/{userId}`).Handler(kithttp.NewServer(
		endpoint,
		decodeDiscoverRequest,
		encodeDiscoverResponse,
	))

	return r
}

type GetUserNameRequest struct {
	UserId string `json:"userId"`
}

type GetUserNameResponse struct {
	UserName string `json:"userName"`
}

func decodeDiscoverRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["userId"]
	if !ok {
		return nil, ErrorBadRequest
	}

	return GetUserNameRequest{
		UserId: userId,
	}, nil
}

func encodeDiscoverResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
