package protocol

import (
	"bufio"
	"errors"
	"strconv"
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

	var body []byte
	if headers["Content-Length"] != "" {
		contentLength, err := strconv.Atoi(headers["Content-Length"])
		if err != nil {
			return nil, errors.New("Invalid Content-Length header: " + err.Error())
		}

		if contentLength > 0 {
			body = make([]byte, contentLength)
			_, err := reader.Read(body)
			if err != nil {
				return nil, errors.New("Error reading body: " + err.Error())
			}
		}
	}

	return &Request{
		method:  statusLine[0],
		path:    statusLine[1],
		version: statusLine[2],
		headers: headers,
		body:    string(body),
	}, nil
}
