package middleware

import (
	"cells/cell-limit/util"
	"context"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"time"
)

func NewRateLimit(interval int, burst int) endpoint.Middleware {
	limiter := rate.NewLimiter(rate.Every(time.Second*time.Duration(interval)), burst)
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limiter.Allow() {
				//这里使用了上节的自定义错误
				return nil, util.NewMyError(429, "too many request，please waiting...")
			}
			return next(ctx, request)
		}
	}
}
