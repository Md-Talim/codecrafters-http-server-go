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

	// Concurrent connections
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

// handleConnection manages a single client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		fmt.Printf("Handling connection from %s\n", conn.RemoteAddr())

		reader := bufio.NewReader(conn)
		req, err := protocol.ParseRequest(reader)
		if err != nil {
			fmt.Println("Error parsing request:", err)
			return
		}

		handler := handlers.GetHandler(req.GetPath())
		res := handler.Handle(req)
		protocol.WriteResponse(conn, res)
	}
}
