package handlers

import (
	"strconv"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

type EchoHandler struct{}

func (e *EchoHandler) Handle(req *protocol.Request) *protocol.Response {
	path := req.GetPath()
	str := path[len("/echo/"):]

	res := &protocol.Response{
		Status:  "200 OK",
		Version: protocol.Version,
		Headers: make(map[string]string),
		Body:    str,
	}

	res.Headers["Content-Type"] = "text/plain"
	res.Headers["Content-Length"] = strconv.Itoa(len(str))

	return res
}
