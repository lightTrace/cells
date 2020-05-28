package main

import (
	"cells/cell-error/userservice"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svc userservice.IUserService
	svc = userservice.UserService{}
	getUserNameEndpoint := userservice.MakeGetUserNameEndpoint(svc)
	updateUserNameEndpoint := userservice.MakeUpdateUserNameEndpoint(svc)
	endpoints := userservice.Endpoints{
		GetUserNameEndpoint:    getUserNameEndpoint,
		UpdateUserNameEndpoint: updateUserNameEndpoint,
	}
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := userservice.MakeHttpHandler(ctx, endpoints, logger)

	go func() {
		fmt.Println("Http Server start at port:8000")
		handler := r
		errChan <- http.ListenAndServe(":8000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
