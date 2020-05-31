package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	kithttp "github.com/go-kit/kit/transport/http"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func userFactory(_ context.Context, method, path string) sd.Factory {
	return func(instance string) (endpoint endpoint.Endpoint, closer io.Closer, err error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = path

		var (
			enc kithttp.EncodeRequestFunc
			dec kithttp.DecodeResponseFunc
		)
		enc, dec = encodeUserRequest, decodeUserResponse

		return kithttp.NewClient(method, tgt, enc, dec).Endpoint(), nil, nil
	}
}

func encodeUserRequest(_ context.Context, req *http.Request, request interface{}) error {
	getUserNameReq := request.(GetUserNameRequest)
	p := "/" + getUserNameReq.UserId
	req.URL.Path += p
	return nil
}

func decodeUserResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response GetUserNameResponse
	var s map[string]interface{}

	if respCode := resp.StatusCode; respCode >= 400 {
		if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
			return nil, err
		}
		return nil, errors.New(s["error"].(string) + "\n")
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
