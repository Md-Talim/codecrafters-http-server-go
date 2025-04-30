package handlers

import (
	"strconv"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

type UserAgentHandler struct{}

func (r *UserAgentHandler) Handle(req *protocol.Request) *protocol.Response {
	userAgent := req.GetHeader("User-Agent")

	res := &protocol.Response{
		Status:  "200 OK",
		Version: protocol.Version,
		Headers: make(map[string]string),
		Body:    userAgent,
	}

	res.Headers["Content-Type"] = "text/plain"
	res.Headers["Content-Length"] = strconv.Itoa(len(userAgent))

	return res
}
