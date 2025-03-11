package server

import (
	"net/http"
	"strconv"
)

// request
type baseRequest struct {
	Content string `json:"content"`
}

type AbstractRequest struct {
	baseRequest
}

type PolishRequest struct {
	baseRequest
}

// response
type Response struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func newSuccessResponse(msg string, data any) Response {
	return Response{
		Code: strconv.Itoa(http.StatusOK),
		Msg:  msg,
		Data: data,
	}
}

type ContentResponse struct {
	Content string `json:"content"`
}
