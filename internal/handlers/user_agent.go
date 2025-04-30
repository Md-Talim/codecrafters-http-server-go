package handlers

import (
	"strconv"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

type UserAgentHandler struct{}

func (r *UserAgentHandler) Handle(req *protocol.Request) *protocol.Response {
	userAgent := req.GetHeader(protocol.HeaderUserAgent)

	res := &protocol.Response{
		Status:  protocol.StatusOK,
		Version: protocol.Version,
		Headers: make(map[string]string),
		Body:    userAgent,
	}

	res.Headers[protocol.HeaderContentType] = protocol.ContentTypeTextPlain
	res.Headers[protocol.HeaderContentLength] = strconv.Itoa(len(userAgent))

	return res
}
