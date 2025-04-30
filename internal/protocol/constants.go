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
	HeaderContentLength   = "Content-Length"
	HeaderContentType     = "Content-Type"
	HeaderUserAgent       = "User-Agent"
	HeaderAcceptEncoding  = "Accept-Encoding"
	HeaderContentEncoding = "Content-Encoding"
	HeaderConnection      = "Connection"
)

// Common Content Types
const (
	ContentTypeTextPlain   = "text/plain"
	ContentTypeOctetStream = "application/octet-stream"
)

// Suppored Header Values
const (
	ContentEncodingGzip = "gzip"
	ConnectionClose     = "close"
)

// HTTP Version
const Version = "HTTP/1.1"
