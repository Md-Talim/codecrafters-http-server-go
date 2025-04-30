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
		Status:  protocol.StatusOK,
		Version: protocol.Version,
		Headers: make(map[string]string),
		Body:    str,
	}

	res.Headers[protocol.HeaderContentType] = protocol.ContentTypeTextPlain
	res.Headers[protocol.HeaderContentLength] = strconv.Itoa(len(str))

	return res
}
