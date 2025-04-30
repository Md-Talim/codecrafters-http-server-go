package protocol

import (
	"bufio"
	"errors"
	"strings"
)

// ParseRequest reads an HTTP request from the given reader and returns a Request object.
// It reads the request line, headers, and the body (based on Content-Length).
// Returns an error if the request is malformed or an I/O error occurs.
func ParseRequest(reader *bufio.Reader) (*Request, error) {
	// Read status line
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, errors.New("Error reading status line: " + err.Error())
	}
	statusLine := strings.SplitN(line, " ", 3)

	// Read headers
	headers := make(Headers)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, errors.New("Error reading headers: " + err.Error())
		}
		if line == "\r\n" {
			break
		}
		header := strings.SplitN(line, ":", 2)
		headers[strings.TrimSpace(header[0])] = strings.TrimSpace(header[1])
	}

	// TODO: Check for Content-Lenght header
	// If Content-Length header exists and > 0,
	// read exactly that many bytes for the body

	return &Request{
		method:  statusLine[0],
		path:    statusLine[1],
		version: statusLine[2],
		headers: headers,
	}, nil
}
