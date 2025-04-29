package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/md-talim/codecrafters-http-server-go/internal/handlers"
	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	handleConnection(conn)
}

// handleConnection manages a single client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Handling connection from %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	req, err := protocol.ParseRequest(reader)
	if err != nil {
		fmt.Println("Error parsing request:", err)
		return
	}

	handler := handlers.GetHandler(req.GetPath())
	if handler == nil {
		return
	}

	res := handler.Handle(req)
	err = protocol.WriteResponse(conn, res)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
