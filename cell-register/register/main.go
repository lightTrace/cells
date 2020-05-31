package main

import (
	"cells/cell-register/register/client"
	"cells/cell-register/register/middleware"
	"cells/cell-register/register/userservice"
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		consulHost  = flag.String("consulHost", "", "consul ip address")
		consulPort  = flag.String("consulPort", "", "consul port")
		serviceHost = flag.String("serviceHost", "", "service ip address")
		servicePort = flag.String("servicePort", "", "service port")
	)
	flag.Parse()

	ctx := context.Background()
	errChan := make(chan error)

	var svc userservice.IUserService
	svc = userservice.UserService{}

	getUserNameEndpoint := userservice.MakeGetUserNameEndpoint(svc)
	//每秒钟getUserName接口只能接受一个请求，但是可以容忍瞬间提高的5个请求，超过限制的请求会报429
	getUserNameEndpoint = middleware.NewRateLimit(1, 5)(getUserNameEndpoint)

	//创建健康检查的Endpoint，未增加限流
	healthEndpoint := userservice.MakeHealthCheckEndpoint(svc)

	updateUserNameEndpoint := userservice.MakeUpdateUserNameEndpoint(svc)
	endpoints := userservice.Endpoints{
		GetUserNameEndpoint:    getUserNameEndpoint,
		UpdateUserNameEndpoint: updateUserNameEndpoint,
		HealthCheckEndpoint:    healthEndpoint,
	}
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := userservice.MakeHttpHandler(ctx, endpoints, logger)

	//创建注册对象
	register := client.Register(*consulHost, *consulPort, *serviceHost, *servicePort, logger)

	go func() {
		fmt.Println("Http Server start at port:" + *servicePort)
		//启动前执行注册
		register.Register()
		handler := r
		errChan <- http.ListenAndServe(":"+*servicePort, handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	//服务退出取消注册
	register.Deregister()
	fmt.Println(error)
}
