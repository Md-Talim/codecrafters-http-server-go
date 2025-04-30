package protocol

import (
	"slices"
	"strings"
)

type Headers map[string]string

type Request struct {
	method  string
	path    string
	version string
	headers Headers
	body    string
}

func (req *Request) GetBody() string {
	return req.body
}

func (req *Request) GetHeader(name string) string {
	if header, ok := req.headers[name]; ok {
		return header
	}
	return ""
}

func (req *Request) GetMethod() string {
	return req.method
}

func (req *Request) GetPath() string {
	return req.path
}

func (req *Request) IsGzipAccepted() bool {
	if header, ok := req.headers[HeaderAcceptEncoding]; ok {
		schemes := strings.Split(header, ", ")
		if slices.Contains(schemes, ContentEncodingGzip) {
			return true
		}
	}
	return false
}

type Response struct {
	Status  string
	Version string
	Headers Headers
	Body    string
}

func (res *Response) getStatusLine() string {
	return res.Version + " " + res.Status
}
