package handlers

import (
	"strings"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

type RequestHandler interface {
	Handle(req *protocol.Request) *protocol.Response
}

var router = map[string]RequestHandler{
	"/":           &RootHandler{},
	"/echo/":      &EchoHandler{},
	"/user-agent": &UserAgentHandler{},
	"/files/":     &FilesHandler{},
}

func GetHandler(path string) RequestHandler {
	if path == "/" {
		return &RootHandler{}
	}
	for p, handler := range router {
		if p == "/" {
			continue
		}
		if strings.HasPrefix(path, p) {
			return handler
		}
	}
	return &NotFoundHandler{}
}
