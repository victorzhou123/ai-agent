package server

type baseRequest struct {
	Content string `json:"content"`
}

type AbstractRequest struct {
	baseRequest
}

type PolishRequest struct {
	baseRequest
}
