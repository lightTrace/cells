package middleware

import (
	"cells/cell-trace/register/userservice"
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

type tracerMiddlewareServer struct {
	userservice.IUserService
	tracer opentracing.Tracer
}

func NewTracerMiddleware(tracer opentracing.Tracer) userservice.ServiceMiddleware {
	return func(next userservice.IUserService) userservice.IUserService {
		return tracerMiddlewareServer{
			next,
			tracer,
		}
	}
}

func (tm tracerMiddlewareServer) GetUserName(ctx context.Context, userId string) (string, error) {
	span, ctxContext := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "service", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("userId", userId)
		span.Finish()
	}()
	return tm.IUserService.GetUserName(ctxContext, userId)
}

func (tm tracerMiddlewareServer) UpdateUserName(ctx context.Context, userId, userName string) error {
	span, _ := opentracing.StartSpanFromContextWithTracer(ctx, tm.tracer, "service", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServerMiddleware",
	})
	defer func() {
		span.LogKV("userId", userId, "userName", userName)
		span.Finish()
	}()
	return tm.IUserService.UpdateUserName(ctx, userId, userName)
}

func NewJaegerTracer(serviceName string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},

		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},

		ServiceName: serviceName,
	}

	tracer, closer, err = cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}
