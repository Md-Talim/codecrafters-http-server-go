package handlers

import (
	"strings"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

type RequestHandler interface {
	Handle(req *protocol.Request) *protocol.Response
}

var routes = map[string]RequestHandler{
	"/":           &RootHandler{},
	"/echo/":      &EchoHandler{},
	"/user-agent": &UserAgentHandler{},
}

func GetHandler(path string) RequestHandler {
	if path == "/" {
		return &RootHandler{}
	}
	for p, handler := range routes {
		if p == "/" {
			continue
		}
		if strings.HasPrefix(path, p) {
			return handler
		}
	}
	return &NotFoundHandler{}
}
