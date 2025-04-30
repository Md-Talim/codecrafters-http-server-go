package handlers

import "github.com/md-talim/codecrafters-http-server-go/internal/protocol"

type NotFoundHandler struct{}

func (r *NotFoundHandler) Handle(req *protocol.Request) *protocol.Response {
	return &protocol.Response{
		Status:  protocol.StatusNotFound,
		Version: protocol.Version,
		Headers: nil,
		Body:    "Not Found",
	}
}
