package protocol

// HTTP Methods
const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

// HTTP Status Codes
const (
	StatusOK       = "200 OK"
	StatusCreated  = "201 Created"
	StatusNotFound = "404 Not Found"
)

// HTTP Header Keys
const (
	HeaderContentLength = "Content-Length"
	HeaderContentType   = "Content-Type"
	HeaderUserAgent     = "User-Agent"
)

// Common Content Types
const (
	ContentTypeTextPlain   = "text/plain"
	ContentTypeOctetStream = "application/octet-stream"
)

// HTTP Version
var Version = "HTTP/1.1"
