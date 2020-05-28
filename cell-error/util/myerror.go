package util

import (
	"context"
	"net/http"
)

type MyError struct {
	Code    int
	Message string
}

func NewMyError(code int, message string) error {
	return &MyError{Code: code, Message: message}
}

func (e *MyError) Error() string {
	return e.Message
}

//自定义一个解码error函数
func MyErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain,charset=utf-8", []byte(err.Error())
	w.Header().Set("content-type", contentType)
	if err, ok := err.(*MyError); ok {
		w.WriteHeader(err.Code)
	} else {
		w.WriteHeader(500)
	}
	w.Write(body)
}
