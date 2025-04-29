package handlers

import "github.com/md-talim/codecrafters-http-server-go/internal/protocol"

type RootHandler struct{}

func (r *RootHandler) Handle(req *protocol.Request) *protocol.Response {
	return &protocol.Response{
		Status:  "200 OK",
		Version: protocol.Version,
		Headers: nil,
		Body:    "",
	}
}
