package handlers

import "github.com/md-talim/codecrafters-http-server-go/internal/protocol"

type RequestHandler interface {
	Handle(req *protocol.Request) *protocol.Response
}

var routes = map[string]RequestHandler{
	"/": &RootHandler{},
}

func GetHandler(path string) RequestHandler {
	handler, ok := routes[path]
	if !ok {
		return nil
	}
	return handler
}
