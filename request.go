package creem

import "net/http"

type Request interface {
	Gateway() string
	Method() string
}

type GetRequest struct{}

func (r GetRequest) Method() string {
	return http.MethodGet
}

type PostRequest struct{}

func (r PostRequest) Method() string {
	return http.MethodPost
}

type DeleteRequest struct{}

func (r DeleteRequest) Method() string {
	return http.MethodDelete
}
