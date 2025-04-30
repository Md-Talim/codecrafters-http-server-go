package handlers

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

// FilesHandler handles HTTP requests for file operations.
// It supports:
//   - GET /files/{filename}: Retrieves the content of the specified file
//   - POST /files/{filename}: Creates a new file with the request body content
type FilesHandler struct{}

func (r *FilesHandler) Handle(req *protocol.Request) *protocol.Response {
	path := req.GetPath()
	method := req.GetMethod()
	body := req.GetBody()
	filename := path[len("/files/"):]

	notFoundHandler := &NotFoundHandler{}

	if method == "GET" {
		res, err := getFile(filename)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return notFoundHandler.Handle(req)
		}
		return res
	} else if method == "POST" {
		res, err := createFile(filename, body)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return notFoundHandler.Handle(req)
		}
		return res
	} else {
		return notFoundHandler.Handle(req)
	}
}

// getFile retrieves the content of a file from the server's file directory
// and returns it as an HTTP response.
//
// Parameters:
//   - filename: The name of the file to retrieve.
//
// Returns:
//   - *protocol.Response: An HTTP response containing the file's content,
//     with appropriate headers such as "Content-Type" and "Content-Length".
//   - error: An error if the file cannot be read or if the directory cannot
//     be determined.
func getFile(filename string) (*protocol.Response, error) {
	directory, err := getDirectory()
	if err != nil {
		return nil, err
	}

	filepath := directory + filename

	// If the file doesn't exist in the files directory
	// return a 404 Not Found response
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	res := &protocol.Response{
		Status:  "200 OK",
		Version: protocol.Version,
		Headers: make(map[string]string),
		Body:    string(file),
	}

	res.Headers["Content-Type"] = "application/octet-stream"
	res.Headers["Content-Length"] = strconv.Itoa(len(file))

	return res, nil
}

// createfile creates a new file with the specified filename and content.
// It retrieves the directory path using getDirectory, constructs the full
// file path, and writes the content to the file. The function returns a
// pointer to a protocol.Response with a "201 Created" status if successful,
// or an error if any step fails.
//
// Parameters:
//   - filename: The name of the file to be created.
//   - content: The content to be written to the file.
//
// Returns:
//   - *protocol.Response: A response object indicating the success of the operation.
//   - error: An error object if the operation fails.
func createFile(filename string, content string) (*protocol.Response, error) {
	directory, err := getDirectory()
	if err != nil {
		return nil, err
	}

	filepath := directory + filename

	err = os.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		return nil, err
	}

	res := &protocol.Response{
		Status:  "201 Created",
		Version: protocol.Version,
		Headers: make(map[string]string),
	}

	return res, nil
}

// getDirectory parses command line arguments to find and return the directory path
// specified by the --directory flag. The directory is used as the root location
// for all file operations.
//
// Returns:
//   - string: The directory path if found
//   - error: An error if the --directory flag is missing or invalid
func getDirectory() (string, error) {
	var directory string
	args := os.Args
	if len(args) <= 2 {
		return "", errors.New("not enough arguments")
	}

	for i, arg := range args {
		if arg == "--directory" {
			if i+1 < len(args) {
				directory = args[i+1]
			} else {
				return "", errors.New("missing directory path")
			}
		}
	}

	return directory, nil
}
