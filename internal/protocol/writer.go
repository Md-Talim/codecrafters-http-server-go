package protocol

import (
	"io"
)

// WriteResponse formats and writes the given Response object to the writer.
// It automatically calculates and adds the Content-Length header if not already present
// and if the response body is not empty.
func WriteResponse(writer io.Writer, res *Response) {
	writer.Write([]byte(res.getStatusLine()))
	writer.Write([]byte("\r\n"))

	for k, v := range res.Headers {
		writer.Write([]byte(k))
		writer.Write([]byte(": "))
		writer.Write([]byte(v))
		writer.Write([]byte("\r\n"))
	}

	writer.Write([]byte("\r\n"))
	writer.Write([]byte(res.Body))
}
