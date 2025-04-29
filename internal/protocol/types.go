package protocol

type Headers map[string]string

var Version = "HTTP/1.1"

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
	header, ok := req.headers[name]
	if !ok {
		return ""
	}
	return header
}

func (req *Request) GetMethod() string {
	return req.method
}

func (req *Request) GetPath() string {
	return req.path
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
