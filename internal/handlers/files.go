package handlers

import (
	"os"
	"strconv"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

// FilesHandler implements the handler for /files/{filename} endpoint,
// which returns a requested file to the client
type FilesHandler struct{}

func (r *FilesHandler) Handle(req *protocol.Request) *protocol.Response {
	// directory stores the value of --directory flag from the args
	// The --directory flag specifies the directory where the files are stored
	var directory string
	args := os.Args
	if len(args) <= 2 {
		return nil
	}
	if args[1] == "--directory" {
		directory = args[2]
	}

	path := req.GetPath()

	// Extract the filename from /files/{filename} endpoint
	filename := path[len("/files/"):]
	filepath := directory + filename

	// If the file doesn't exist in the files directory
	// return a 404 Not Found response
	file, err := os.ReadFile(filepath)
	if err != nil {
		handler := &NotFoundHandler{}
		return handler.Handle(req)
	}

	res := &protocol.Response{
		Status:  "200 OK",
		Version: protocol.Version,
		Headers: make(map[string]string),
		Body:    string(file),
	}

	res.Headers["Content-Type"] = "application/octet-stream"
	res.Headers["Content-Length"] = strconv.Itoa(len(file))

	return res
}
