package handlers

import (
	"bytes"
	"compress/gzip"
	"strconv"

	"github.com/md-talim/codecrafters-http-server-go/internal/protocol"
)

type EchoHandler struct{}

func (e *EchoHandler) Handle(req *protocol.Request) *protocol.Response {
	path := req.GetPath()
	str := path[len("/echo/"):]

	res := &protocol.Response{
		Status:  protocol.StatusOK,
		Version: protocol.Version,
		Headers: make(map[string]string),
	}

	res.Headers[protocol.HeaderContentType] = protocol.ContentTypeTextPlain
	res.Headers[protocol.HeaderContentLength] = strconv.Itoa(len(str))

	if _, ok := req.AcceptEncodingHeader(); ok {
		compressedBody := compressData(str)
		res.Body = string(compressedBody)
		res.Headers[protocol.HeaderContentEncoding] = protocol.ContentEncodingGzip
		res.Headers[protocol.HeaderContentLength] = strconv.Itoa(len(compressedBody))
	} else {
		res.Body = str
	}

	return res
}

func compressData(data string) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err := gz.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	gz.Close()
	return b.Bytes()
}
